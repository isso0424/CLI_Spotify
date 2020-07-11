package command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

  "isso0424/spotify-rapspi/types"
)

func GetPlayStatus(token string) {
  response, err := createRequest(token, "GET", "https://api.spotify.com/v1/me/player")
  if err != nil {
    fmt.Println(err)
    return
  }
  buffer := make([]byte, 8192)
  _, err = response.Body.Read(buffer)

  buffer = bytes.Trim(buffer, "\x00")

  var responseBody types.Content
  if err := json.Unmarshal(buffer, &responseBody); err != nil {
    log.Fatal(err)
  }

  createInfo(responseBody)
}

func createInfo(content types.Content) {
  if content.IsPlaying {
    fmt.Printf("Playing status\n--------------\nTitle: %s\nArtist: %s\n", content.Item.Name, content.Item.Artists[0].Name)
  } else {
    fmt.Println("Pausing")
  }
}
