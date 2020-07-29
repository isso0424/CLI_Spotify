package command

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"isso0424/spotify_CLI/command/file"
	"isso0424/spotify_CLI/command/request"
	"isso0424/spotify_CLI/selfmadetypes"
	"isso0424/spotify_CLI/util"
	"strconv"
	"strings"
)

const (
	track = "track"
	off   = "off"
)

func playFromURL(token *string, uri string) (err error) {
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

	_, statusCode, err := request.CreateRequest(token, selfmadetypes.PUT, "/me/player/play", bytes.NewBuffer(values))
	if err != nil {
		return
	}

	fmt.Println(statusCode)

	err = status{}.Execute(token)

	if err != nil {
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

func switchRepeatState(state string) string {
	switch state {
	case track:
		return off
	case "context":
		return track
	case off:
		return "context"
	}

	return off
}

// Execute is excution command function.
func (cmd welcome) Execute(token *string) (err error) {
	response, _, err := request.CreateRequest(token, selfmadetypes.GET, "/me", nil)
	if err != nil {
		return
	}

	var userInfo selfmadetypes.User
	err = json.Unmarshal(response, &userInfo)
	if err != nil {
		return
	}

	fmt.Printf("ようこそ! %sさん!\n", userInfo.DisplayName)

	return
}

func saveSearchResult(searchResults []selfmadetypes.SearchResultItem) (err error) {
	var isSave string
	util.Input("Want to save result?\n------------------------", "Want to save?", &isSave)

	if isSave != "yes" {
		return
	}

	var rawIndex string
	util.Input("Please input index\n------------------------", "Index", &rawIndex)

	index, err := strconv.Atoi(rawIndex)
	if err != nil {
		return
	}

	if index >= len(searchResults) {
		return errors.New("index is out of range")
	}

	item := searchResults[index]

	err = file.SavePlayList(selfmadetypes.PlayList{Name: item.Name, URI: item.URI})

	return
}

func existTarget(target string, judgeTargets []string) bool {
	for _, judgeTarget := range judgeTargets {
		if judgeTarget == target {
			return true
		}
	}

	return false
}

func getArtistsName(artists []selfmadetypes.Artists) (artistNames string) {
	artistNames = ""

	for _, artist := range artists {
		artistNames += artist.Name + " "
	}

	return
}

func getPlayingPlaylistID(token *string) (id *string, err error) {
	playingStatus, err := request.GetStatus(token)
	if err != nil {
		return
	}

	url := playingStatus.Context.Href

	id = &strings.Split(url, "/")[5]

	return
}
