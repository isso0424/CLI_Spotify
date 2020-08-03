package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"isso0424/spotify_CLI/command/parse"
	"isso0424/spotify_CLI/selfmadetypes"
	"strings"
)

const track = "track"

// PlayFromURL is play track or playlist.
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

	_, err = CreateRequest(token, selfmadetypes.PUT, "/me/player/play", bytes.NewBuffer(values))
	if err != nil {
		return
	}

	err = PrintPlayingStatus(token)

	return
}

// PrintPlayingStatus is function that print playing status.
func PrintPlayingStatus(token *string) (err error) {
	status, err := GetStatus(token)
	isPause := err == io.EOF || (err == nil && status == nil)
	if isPause {
		print("Pausing")
		return
	}
	if err != nil {
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
