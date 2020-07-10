package command

import (
	"fmt"
	"net/http"
	"net/url"

	"golang.org/x/oauth2"
)

func Resume(token *oauth2.Token) {
  values := url.Values{}
  values.Add("context_uri", "spotify:playlist:4Cu0w3ZsPYiPQTeElf2Tls")
  request, err := http.NewRequest("PUT", "https://api.spotify.com/v1/me/player/play", nil)
  if err != nil {
    fmt.Printf(err.Error())
  }

  request.Header.Set("Authorization", "Bearer " + token.AccessToken)
  client := &http.Client{}
  response, err := client.Do(request)
  if err != nil {
    fmt.Printf(err.Error())
    return
  }
  defer response.Body.Close()
}
