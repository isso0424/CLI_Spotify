package main

import (
	"fmt"
	"log"
	"net/http"

  "isso0424/spotify-rapspi/command"
	"github.com/zmb3/spotify"
)

var (
  auth = spotify.NewAuthenticator("http://localhost:8888/callback", spotify.ScopeUserModifyPlaybackState, spotify.ScopeUserReadPlaybackState)
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
  command.GetPlayStatus(token)
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
