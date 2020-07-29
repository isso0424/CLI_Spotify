package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"isso0424/spotify_CLI/command/parse"
	"isso0424/spotify_CLI/selfmadetypes"
	"strings"
)

const track = "track"

func PlayFromURL(token *string, uri string) (err error) {
	uriKind := strings.Split(uri, ":")[1]
	var values []byte
	if uriKind == track {
		values, err = json.Marshal(playJSON{URIs: []string{uri}})
	} else {
		values, err = json.Marshal(playListJSON{ContextURI: uri})
	}
	if err != nil {
		return
	}
	fmt.Println(string(values))

	_, _, err = CreateRequest(token, selfmadetypes.PUT, "/me/player/play", bytes.NewBuffer(values))
	if err != nil {
		return
	}

	err = PrintPlayingStatus(token)

	return
}

func PrintPlayingStatus(token *string) (err error) {
	status, err := GetStatus(token)
	if err != nil || status == nil {
		return
	}

	url := status.Context.ExternalUrls.Spotify
	id, err := parse.GetPlaylistID(url)
	if err != nil {
		return
	}

	listStatus, err := GetPlayListStatus(token, id)
	if err != nil {
		return
	}

	fmt.Println(parse.CreatePlayingStatus(*status, listStatus))

	return
}

type playListJSON struct {
	ContextURI string `json:"context_uri"`
}

type playJSON struct {
	URIs []string `json:"uris"`
}
