package command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

  "isso0424/spotify-rapspi/types"
)

func GetPlayStatus(token string) bool {
  request, err := http.NewRequest("GET", "https://api.spotify.com/v1/me/player", nil)
  if err != nil {
    fmt.Printf(err.Error())
  }

  request.Header.Set("Authorization", "Bearer " + token)
  client := &http.Client{}
  response, err := client.Do(request)
  if err != nil {
    fmt.Printf(err.Error())
    return false
  }
  buffer := make([]byte, 8192)
  _, err = response.Body.Read(buffer)

  buffer = bytes.Trim(buffer, "\x00")

  var responseBody types.Content
  if err := json.Unmarshal(buffer, &responseBody); err != nil {
    log.Fatal(err)
  }


  defer response.Body.Close()

  return responseBody.IsPlaying
}
