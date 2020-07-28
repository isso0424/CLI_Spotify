package command

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"isso0424/spotify_CLI/auth"
	"isso0424/spotify_CLI/command/file"
	"isso0424/spotify_CLI/command/parse"
	"isso0424/spotify_CLI/command/request"
	"isso0424/spotify_CLI/selfmadetypes"
	"isso0424/spotify_CLI/util"
	"math/rand"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	track = "track"
	off   = "off"
)

// Execute is excution command function.
func (cmd status) Execute(token *string) error {
	status, err := request.GetStatus(token)
	if err != nil {
		return err
	}

	if status == nil {
		return nil
	}
	playlistURL := status.Context.ExternalUrls.Spotify
	playlistID, err := parse.GetPlaylistID(playlistURL)

	if err != nil {
		return err
	}

	playListStatus, err := request.GetPlayListStatus(token, playlistID)

	if err != nil {
		return err
	}

	fmt.Println(parse.CreatePlayingStatus(*status, playListStatus))

	return nil
}

// Execute is excution command function.
func (cmd next) Execute(token *string) (err error) {
	_, _, err = request.CreateRequest(token, selfmadetypes.POST, "/me/player/next", nil)

	if err != nil {
		return
	}

	err = status{}.Execute(token)

	return
}

// Execute is excution command function.
func (cmd pause) Execute(token *string) (err error) {
	_, _, err = request.CreateRequest(token, selfmadetypes.PUT, "/me/player/pause", nil)

	if err != nil {
		return
	}
	fmt.Println("paused!!!")

	return
}

// Execute is excution command function.
func (cmd play) Execute(token *string) (err error) {
	var url string
	util.Input("please input playlist url\n------------------------", "PlayListURL", &url)

	uri, err := parse.CreateContextURI(url)
	if err != nil {
		return
	}
	err = playFromURL(token, *uri)

	return
}

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

// Execute is excution command function.
func (cmd prev) Execute(token *string) (err error) {
	_, _, err = request.CreateRequest(token, selfmadetypes.POST, "/me/player/previous", nil)

	if err != nil {
		return
	}

	err = status{}.Execute(token)

	return
}

type playListJSON struct {
	ContextURI string `json:"context_uri"`
}

type playJSON struct {
	URIs []string `json:"uris"`
}

func choice(playlists []selfmadetypes.PlayList) selfmadetypes.PlayList {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(playlists))

	return playlists[index]
}

// Execute is excution command function.
func (cmd repeat) Execute(token *string) (err error) {
	status, err := request.GetStatus(token)

	if err != nil {
		return
	}

	state := switchRepeatState(status.RepeatState)

	_, _, err = request.CreateRequest(token, selfmadetypes.PUT, fmt.Sprintf("/me/player/repeat?state=%s", state), nil)

	if err != nil {
		return
	}

	fmt.Printf("Repeat state change to `%s`\n", state)

	return
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
func (cmd resume) Execute(token *string) (err error) {
	_, _, err = request.CreateRequest(token, selfmadetypes.PUT, "/me/player/play", nil)

	if err != nil {
		return
	}
	fmt.Println("resumed!!!")

	return
}

// Execute is excution command function.
func (cmd shuffle) Execute(token *string) (err error) {
	status, err := request.GetStatus(token)
	if err != nil {
		return
	}

	state := !status.ShuffleState

	_, _, err = request.CreateRequest(token, selfmadetypes.PUT, fmt.Sprintf("/me/player/shuffle?state=%v", state), nil)
	if err != nil {
		return
	}

	fmt.Printf("Switch shuffle state to %v\n", state)

	return
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

// Execute is excution command function.
func (cmd refresh) Execute(token *string) error {
	tokenPtr, err := auth.GetToken()
	if err != nil {
		return err
	}

	*token = *tokenPtr

	return nil
}

// Execute is excution command function.
func (cmd volume) Execute(token *string) (err error) {
	var percent string
	util.Input("please volume percent\n------------------------", "Volume", &percent)

	percentInt, err := strconv.Atoi(percent)
	if err != nil {
		return
	}

	if percentInt < 0 || percentInt > 100 {
		return errors.New("percent range is 0 to 100")
	}

	_, _, err = request.CreateRequest(
		token,
		selfmadetypes.PUT,
		fmt.Sprintf(
			"/me/player/volume?volume_percent=%s",
			percent,
		),
		nil,
	)

	return
}

// Execute is excution command function.
func (cmd search) Execute(token *string) (err error) {
	var kind string
	util.Input(
		"please input search kind\n\n"+
			"search kinds: album artist playlist track show episode\n\n"+
			"if input over 2 types, please enter with a colon\n"+
			"------------------------",
		"Kind",
		&kind,
	)
	kinds := strings.Split(kind, ",")
	for _, kind := range kinds {
		if existTarget(kind, []string{"album", "artist", "playlist", "track", "show", "episode"}) {
			return fmt.Errorf("search type %s is not found", kind)
		}
	}

	var keyword string
	util.Input("Please input search keyword\n------------------------", "Keyword", &keyword)
	keyword = url.QueryEscape(keyword)

	response, _, err := request.CreateRequest(
		token,
		selfmadetypes.GET,
		fmt.Sprintf(
			"/search?q=%s&type=%s",
			keyword,
			kind,
		),
		nil,
	)
	if err != nil {
		return
	}

	var searchResponse selfmadetypes.SearchResponse
	err = json.Unmarshal(response, &searchResponse)
	if err != nil {
		return
	}

	searchResultItems := searchResponse.ParseAndPrint(kinds)

	err = saveSearchResult(searchResultItems)

	return err
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

// Execute is excution command function.
func (cmd favoriteTrack) Execute(token *string) (err error) {
	response, _, err := request.CreateRequest(token, selfmadetypes.GET, "/me/player/currently-playing", nil)
	if err != nil {
		return
	}

	var playingStatus selfmadetypes.CurrentPlayStatus

	response = bytes.Trim(response, "\x00")
	err = json.Unmarshal(response, &playingStatus)
	if err != nil {
		return
	}

	id := strings.Split(playingStatus.Item.URI, ":")[2]
	_, _, err = request.CreateRequest(token, selfmadetypes.PUT, fmt.Sprintf("/me/tracks?ids=%s", id), nil)
	if err != nil {
		return
	}

	fmt.Printf("Success add '%s' to your favorite song!!!\n", playingStatus.Item.Name)

	return
}

// Execute is excution command function.
func (cmd importOwnPlaylists) Execute(token *string) (err error) {
	response, _, err := request.CreateRequest(token, selfmadetypes.GET, "/me/playlists", nil)
	if err != nil {
		return
	}

	var userPlayLists selfmadetypes.UserPlaylists
	err = json.Unmarshal(response, &userPlayLists)
	if err != nil {
		return
	}

	for _, playlist := range userPlayLists.Item {
		err = file.SavePlayList(playlist)
		if err != nil {
			return
		}
	}

	return
}

// Execute is excution command function.
func (cmd recent) Execute(token *string) (err error) {
	response, _, err := request.CreateRequest(token, selfmadetypes.GET, "/me/player/recently-played?limit=1", nil)
	if err != nil {
		return
	}

	var recentPlayedTracks selfmadetypes.RecentPlayedTracks
	err = json.Unmarshal(response, &recentPlayedTracks)
	if err != nil {
		return
	}

	recentPlayedTrack := recentPlayedTracks.Items[0]

	fmt.Printf(
		"TrackName: %s\n" +
		"Artist:    %s\n",
		recentPlayedTrack.Name,
		recentPlayedTrack.Artists,
	)

	return
}
