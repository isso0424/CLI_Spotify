package command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func playFromURL(token string, url string) {
  uri, err := createContextUri(url)
  if err != nil {
    log.Fatalln(err)
    return
  }
  values, err := json.Marshal(playJson{ContextUri: *uri})
  request, err := http.NewRequest("PUT", "https://api.spotify.com/v1/me/player/play", bytes.NewBuffer(values))
  if err != nil {
    log.Fatalln(err)
    return
  }

  request.Header.Set("Authorization", "Bearer " + token)
  client := &http.Client{}
  _, err = client.Do(request)
  if err != nil {
    log.Fatalln(err)
    return
  }
}

func createContextUri(url string) (*string, error){
  err := &lengthError{}
  spritted := strings.Split(url, "/")

  if len(spritted) < 5 {
    return nil, err
  }
  kind := spritted[3]
  tmp := spritted[4]

  id := strings.Split(tmp, "?")[0]

  context_uri := fmt.Sprintf("spotify:%s:%s", kind, id)
  return &context_uri, nil
}

func(e *lengthError) Error() string {
  return "too short length"
}

type lengthError struct {
}

type playJson struct {
  ContextUri string `json:"context_uri"`
}
