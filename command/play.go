package command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"isso0424/spotify-rapspi/util"
	"log"
	"net/http"
	"strings"
)

func PlayFromURL(token string) {
  fmt.Printf("please input playlist url\n------------------------")
  var url string
  util.Input("PlayListURL", &url)
  uri, err := CreateContextUri(url)
  if err != nil {
    log.Fatalln(err)
    return
  }
  play(token, *uri)
}

func play(token string, uri string) {
  values, err := json.Marshal(playJson{ContextUri: uri})
  request, err := http.NewRequest("PUT", "https://api.spotify.com/v1/me/player/play", bytes.NewBuffer(values))
  if err != nil {
    log.Fatalln(err)
    return
  }

  request.Header.Set("Authorization", "Bearer " + token)
  client := &http.Client{}
  if _, err = client.Do(request); err != nil {
    log.Fatalln(err)
    return
  }

  if !GetPlayStatus(token) {
    fmt.Println("this url is invalid")
  }
}

func CreateContextUri(url string) (*string, error){
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
