package command

import (
  "fmt"
)

func Pause(token string) {
  _, err := createRequest(token, "PUT", "https://api.spotify.com/v1/me/player/pause")

  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println("paused!!!")
}
