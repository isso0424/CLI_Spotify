package auth

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/zmb3/spotify"
)

var (
  auth = spotify.NewAuthenticator("http://localhost:8888/callback", spotify.ScopeUserModifyPlaybackState, spotify.ScopeUserReadPlaybackState)
  state = "abc123"
  ch = make (chan *spotify.Client)
  tokenFile = ".token"
)

func GetToken() (string, error) {
  if existDotToken() {
    return readDotToken()
  }

  return oauth()
}

func existDotToken() bool {
  _, err := os.Stat(tokenFile)
  return err == nil
}
func readDotToken() (string, error) {
  bytes, err := ioutil.ReadFile(tokenFile)
  if err != nil {
    return "", err
  }

  return string(bytes), nil
}

func oauth() (string, error) {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("error loading .env file")
    return "", err
  }
  auth.SetAuthInfo(os.Getenv("SPOTIFY_ID"), os.Getenv("SPOTIFY_SECRET"))

  http.HandleFunc("/callback", handler)
  http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
    log.Println("Got request for:", request.URL.String())
  })
  go http.ListenAndServe(":8888", nil)

  authUrl := auth.AuthURL(state)
  fmt.Println("Please log in to Spotify by visiting the following page in your browser:", authUrl)

  client := <-ch

  token, err := client.Token()
  if err != nil {
    fmt.Printf(err.Error())
    return "", err
  }

  createDotToken(token.AccessToken)

  return token.AccessToken, nil
}

func createDotToken(token string) error {
  file, err := os.Create(tokenFile)
  if err != nil {
    return err
  }
  file.Close()

  bytes := []byte(token)

  return ioutil.WriteFile(tokenFile, bytes, 0666)
}

func handler(writer http.ResponseWriter, request *http.Request) {
  token, err := auth.Token(state, request)
  if err != nil {
    http.Error(writer, "Couldn't get token", http.StatusForbidden)
    log.Fatal(err)
  }
  if st := request.FormValue("state"); st != state {
    http.NotFound(writer, request)
    log.Fatalf("State mismatch: %s != %s\n", st, state)
  }

  client := auth.NewClient(token)
  fmt.Fprintf(writer, "Login!!!")
  ch <- &client
}
