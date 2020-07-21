package command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"isso0424/spotify_CLI/command/parse"
	"isso0424/spotify_CLI/command/request"
	"isso0424/spotify_CLI/util"
)

func playFromURL(token string) (newToken string, err error) {
	newToken = token

	fmt.Printf("please input playlist url\n------------------------")
	var url string
	util.Input("PlayListURL", &url)

	uri, err := parse.CreateContextUri(url)
	if err != nil {
		return
	}
	newToken, err = play(token, *uri)

	return
}

func play(token string, uri string) (newToken string, err error) {
	newToken = token
	values, err := json.Marshal(playJson{ContextUri: uri})
	if err != nil {
		return
	}

	_, newToken, err = request.CreateRequest(token, "PUT", "https://api.spotify.com/v1/me/player/play", bytes.NewBuffer(values))

  if err != nil {
    return
  }

	nowPlaying, newToken, err := getPlayStatus(token)

	if err != nil {
		return
	}

	if !nowPlaying {
		fmt.Println("this url is invalid")
	}

	return
}

type playJson struct {
	ContextUri string `json:"context_uri"`
}
