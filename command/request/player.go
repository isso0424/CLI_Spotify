package request

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"isso0424/spotify_CLI/command/parse"
	"isso0424/spotify_CLI/selfmadetypes/requesttypes"
	"isso0424/spotify_CLI/util"
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

	_, err = CreateRequest(token, requesttypes.PUT, "/me/player/play", bytes.NewBuffer(values))
	if err != nil {
		return
	}

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
	id, err := parse.GetIDFromURL(url)
	if err != nil {
		return
	}

	kind, err := parse.GetKindFromURL(url)
	if err != nil {
		return
	}

	contextName, contextUser, err := getContextInformation(token, id, kind)
	if err != nil {
		return
	}

	util.Output(parse.CreatePlayingStatus(*status, contextName, contextUser, kind))

	return
}

func getContextInformation(token *string, id, kind string) (name string, user string, err error) {
	switch kind {
	case "playlist":
		listStatus, err := GetPlayListStatus(token, id)
		if err != nil {
			return "", "", err
		}
		name = listStatus.Name
		user = listStatus.Owner.DisplayName
	case "album":
		albumStatus, err := GetAlbumStatus(token, id)
		if err != nil {
			return "", "", err
		}
		name = albumStatus.Name
		user = albumStatus.Artists[0].Name
	case "artist":
		artistStatus, err := GetArtistStatus(token, id)
		if err != nil {
			return "", "", err
		}
		name = artistStatus.Name
	default:
		err = errors.New("kind not found")
		return
	}

	return
}

type playListJSON struct {
	ContextURI string `json:"context_uri"`
}

type playJSON struct {
	URIs []string `json:"uris"`
}
