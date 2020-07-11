package command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

  "isso0424/spotify-rapspi/types"
)

func GetPlayStatus(token string) {
  request, err := http.NewRequest("GET", "https://api.spotify.com/v1/me/player", nil)
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
