package command

import "fmt"

func Prev(token string) (newToken string) {
  _, newToken, err := createRequest(token, "POST", "https://api.spotify.com/v1/me/player/previous", nil)

  if err != nil {
    fmt.Println("Error: ", err)
    return
  }

  _, newToken = GetPlayStatus(token)

  return
}
