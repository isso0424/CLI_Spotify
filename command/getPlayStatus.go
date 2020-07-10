package command

import (
  "fmt"
  "net/http"
	"golang.org/x/oauth2"
)

func GetPlayStatus(token *oauth2.Token) {
  request, err := http.NewRequest("GET", "https://api.spotify.com/v1/me/player", nil)
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
  buffer := make([]byte, 8192)
  size, err := response.Body.Read(buffer)
  fmt.Printf("size: %d\ncontent: %s", size, string(buffer))
  defer response.Body.Close()
}
