package command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"isso0424/spotify_CLI/util"
	"strings"
)

func playFromURL(token string) (newToken string) {
	newToken = token
	fmt.Printf("please input playlist url\n------------------------")
	var url string
	util.Input("PlayListURL", &url)
	uri, err := CreateContextUri(url)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	newToken = play(token, *uri)

	return
}

func play(token string, uri string) (newToken string, err error) {
	newToken = token
	values, err := json.Marshal(playJson{ContextUri: uri})
	if err != nil {
		return
	}

	_, newToken, err = util.CreateRequest(token, "PUT", "https://api.spotify.com/v1/me/player/play", bytes.NewBuffer(values))

	nowPlaying, newToken, err := getPlayStatus(token)

  if err != nil {
    return
  }

	if !nowPlaying {
		fmt.Println("this url is invalid")
	}

	return
}

func CreateContextUri(url string) (*string, error) {
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

func (e *lengthError) Error() string {
	return "too short length"
}

type lengthError struct {
}

type playJson struct {
	ContextUri string `json:"context_uri"`
}
