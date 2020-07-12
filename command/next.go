package command

import "fmt"

func Next(token string) {
  _, err := createRequest(token, "POST", "https://api.spotify.com/v1/me/player/next")

  if err != nil {
    fmt.Println("Error: ", err)
    return
  }

  GetPlayStatus(token)
}
