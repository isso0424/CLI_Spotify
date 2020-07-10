package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
)

var (
  auth = spotify.NewAuthenticator("http://localhost:8888/callback", spotify.ScopeUserModifyPlaybackState)
  state = "abc123"
  ch = make (chan *spotify.Client)
)

func main() {
  http.HandleFunc("/callback", handler)
  http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
    log.Println("Got request for:", request.URL.String())
  })
  go http.ListenAndServe(":8888", nil)

  authUrl := auth.AuthURL(state)
  fmt.Println("Please log in to Spotify by visiting the following page in your browser:", authUrl)

  client := <-ch

  user, err := client.CurrentUser()
  token, err := client.Token()
  if err != nil {
    fmt.Printf(err.Error())
    return
  }
  pause(token)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("you are logged in as:", user.ID)
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

func pause(token *oauth2.Token) {
  request, _ := http.NewRequest("PUT", "https://api.spotify.com/v1/me/player/pause", nil)
  request.Header.Set("Authorization", "Bearer " + token.AccessToken)
  client := &http.Client{}
  response, err := client.Do(request)
  if err != nil {
    fmt.Printf(err.Error())
    return
  }
  defer response.Body.Close()
}
