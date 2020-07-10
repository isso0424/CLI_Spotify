package command

import (
  "fmt"
  "net/http"
	"golang.org/x/oauth2"
)

func Pause(token *oauth2.Token) {
  request, err := http.NewRequest("PUT", "https://api.spotify.com/v1/me/player/pause", nil)
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
