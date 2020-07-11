package command

import (
	"fmt"
	"net/http"
	"net/url"
)

func Resume(token string) {
  values := url.Values{}
  values.Add("context_uri", "spotify:playlist:4Cu0w3ZsPYiPQTeElf2Tls")
  request, err := http.NewRequest("PUT", "https://api.spotify.com/v1/me/player/play", nil)
  if err != nil {
    fmt.Printf(err.Error())
  }

  request.Header.Set("Authorization", "Bearer " + token)
  client := &http.Client{}
  response, err := client.Do(request)
  if err != nil {
    fmt.Printf(err.Error())
    return
  }
  defer response.Body.Close()
}
