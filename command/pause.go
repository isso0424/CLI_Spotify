package command

import (
  "fmt"
  "net/http"
)

func Pause(token string) {
  request, err := http.NewRequest("PUT", "https://api.spotify.com/v1/me/player/pause", nil)
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
