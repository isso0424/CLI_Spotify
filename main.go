package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"isso0424/spotify-rapspi/command"

	"github.com/joho/godotenv"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
)

var (
  auth = spotify.NewAuthenticator("http://localhost:8888/callback", spotify.ScopeUserModifyPlaybackState, spotify.ScopeUserReadPlaybackState)
  state = "abc123"
  ch = make (chan *spotify.Client)
)

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("error loading .env file")
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
    return
  }

  mainLoop(token)
}

func mainLoop(token *oauth2.Token) {
  fmt.Println("if you wanna exit, you must type 'exit'")
  for {
    status := command.GetPlayStatus(token)

    if status {
      command.Pause(token)
      fmt.Println("paused!!!")
    }else {
      command.Resume(token)
      fmt.Println("resumed!!!")
    }

    var input string
    fmt.Scanln(&input)

    if input == "exit" {
      return
    }
  }
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
