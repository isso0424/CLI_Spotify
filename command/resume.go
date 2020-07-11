package command

import (
	"fmt"
)

func Resume(token string) {
  _, err := createRequest(token, "PUT", "https://api.spotify.com/v1/me/player/play")

  if err != nil {
    fmt.Println(err)
  }
}
