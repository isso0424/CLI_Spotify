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
	"isso0424/spotify_CLI/selfMadeTypes"
	"isso0424/spotify_CLI/util"
	"math/rand"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func (_ status) Execute(token *string) error {
	status, err := request.GetStatus(token)
	if err != nil {
		return err
	}

	if status == nil {
		return nil
	}
	playlistUrl := status.Context.ExternalUrls.Spotify
	playlistID, err := parse.GetPlaylistID(playlistUrl)

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

func (_ next) Execute(token *string) (err error) {
	_, err = request.CreateRequest(token, selfMadeTypes.POST, "/me/player/next", nil)

	if err != nil {
		return
	}

	err = status{}.Execute(token)

	return
}

func (cmd pause) Execute(token *string) (err error) {
	_, err = request.CreateRequest(token, selfMadeTypes.PUT, "/me/player/pause", nil)

	if err != nil {
		return
	}
	fmt.Println("paused!!!")

	return
}

func (_ play) Execute(token *string) (err error) {
	var url string
	util.Input("please input playlist url\n------------------------", "PlayListURL", &url)

	uri, err := parse.CreateContextUri(url)
	if err != nil {
		return
	}
	err = playFromURL(token, *uri)

	return
}

func playFromURL(token *string, uri string) (err error) {
	uriKind := strings.Split(uri, ":")[1]
	var values []byte
	if uriKind == "track" {
		values, err = json.Marshal(playJson{Uris: []string{uri}})
	} else {
		values, err = json.Marshal(playListJson{ContextUri: uri})
	}
	if err != nil {
		return
	}
	fmt.Println(string(values))

	response, err := request.CreateRequest(token, selfMadeTypes.PUT, "/me/player/play", bytes.NewBuffer(values))
	if err != nil {
		return
	}

	fmt.Println(response.StatusCode)

	err = status{}.Execute(token)

	if err != nil {
		return
	}

	return
}

func (_ prev) Execute(token *string) (err error) {
	_, err = request.CreateRequest(token, selfMadeTypes.POST, "/me/player/previous", nil)

	if err != nil {
		return
	}

	err = status{}.Execute(token)

	return
}

type playListJson struct {
	ContextUri string `json:"context_uri"`
}

type playJson struct {
	Uris []string `json:"uris"`
}

func choice(playlists []selfMadeTypes.PlayList) selfMadeTypes.PlayList {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(playlists))

	return playlists[index]
}
func (_ repeat) Execute(token *string) (err error) {
	status, err := request.GetStatus(token)

	if err != nil {
		return
	}

	state := switchRepeatState(status.RepeatState)

	_, err = request.CreateRequest(token, selfMadeTypes.PUT, fmt.Sprintf("/me/player/repeat?state=%s", state), nil)

	if err != nil {
		return
	}

	fmt.Printf("Repeat state change to `%s`\n", state)

	return
}

func switchRepeatState(state string) string {
	switch state {
	case "track":
		return "off"
	case "context":
		return "track"
	case "off":
		return "context"
	}

	return "off"
}

func (_ resume) Execute(token *string) (err error) {
	_, err = request.CreateRequest(token, selfMadeTypes.PUT, "/me/player/play", nil)

	if err != nil {
		return
	}
	fmt.Println("resumed!!!")

	return
}

func (_ shuffle) Execute(token *string) (err error) {
	status, err := request.GetStatus(token)
	if err != nil {
		return
	}

	state := !status.ShuffleState

	_, err = request.CreateRequest(token, selfMadeTypes.PUT, fmt.Sprintf("/me/player/shuffle?state=%v", state), nil)
	if err != nil {
		return
	}

	fmt.Printf("Switch shuffle state to %v\n", state)

	return
}

func (_ welcome) Execute(token *string) (err error) {
	response, err := request.CreateRequest(token, selfMadeTypes.GET, "/me", nil)
	if err != nil {
		return
	}

	buffer := make([]byte, 8192)
	_, err = response.Body.Read(buffer)
	if err != nil {
		return
	}

	buffer = bytes.Trim(buffer, "\x00")

	var userInfo selfMadeTypes.User
	err = json.Unmarshal(buffer, &userInfo)
	if err != nil {
		return
	}

	fmt.Printf("ようこそ! %sさん!\n", userInfo.DisplayName)

	return
}

func (_ refresh) Execute(token *string) error {
	tokenPtr, err := auth.GetToken()
	if err != nil {
		return err
	}

	*token = *tokenPtr

	return nil
}

func (_ volume) Execute(token *string) (err error) {
	var percent string
	util.Input("please volume percent\n------------------------", "Volume", &percent)

	percentInt, err := strconv.Atoi(percent)
	if err != nil {
		return
	}

	if percentInt < 0 || percentInt > 100 {
		return errors.New("percent range is 0 to 100")
	}

	_, err = request.CreateRequest(
    token,
    selfMadeTypes.PUT,
    fmt.Sprintf(
    "/me/player/volume?volume_percent=%s",
    percent,
  ),
  nil,
)

	return
}

func (_ search) Execute(token *string) (err error) {
	var kind string
	util.Input(
    "please input search kind\n\n" +
    "search kinds: album artist playlist track show episode\n\n" +
    "if input over 2 types, please enter with a colon\n" +
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

	response, err := request.CreateRequest(
    token,
    selfMadeTypes.GET,
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

	buffer := make([]byte, 65536)
	_, err = response.Body.Read(buffer)
	if err != nil {
		return
	}

	buffer = bytes.Trim(buffer, "\x00")

	var searchResponse selfMadeTypes.SearchResponse
	err = json.Unmarshal(buffer, &searchResponse)
	if err != nil {
		return
	}

	searchResultItems := searchResponse.ParseAndPrint(kinds)

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

	if index >= len(searchResultItems) {
		return errors.New("index is out of range")
	}

	item := searchResultItems[index]

	err = file.SavePlayList(selfMadeTypes.PlayList{Name: item.Name, Uri: item.Uri})

	return
}

func existTarget(target string, judgeTargets []string) bool {
  for _, judgeTarget := range(judgeTargets) {
    if judgeTarget == target {
      return true
    }
  }

  return false
}

func (_ favoriteTrack) Execute(token *string) (err error) {
	response, err := request.CreateRequest(token, selfMadeTypes.GET, "/me/player/currently-playing", nil)
	if err != nil {
		return
	}

	buffer := make([]byte, 65536)
	_, err = response.Body.Read(buffer)
	if err != nil {
		return
	}

	var playingStatus selfMadeTypes.CurrentPlayStatus

	buffer = bytes.Trim(buffer, "\x00")
	err = json.Unmarshal(buffer, &playingStatus)
	if err != nil {
		return
	}

	id := strings.Split(playingStatus.Item.Uri, ":")[2]
	_, err = request.CreateRequest(token, selfMadeTypes.PUT, fmt.Sprintf("/me/tracks?ids=%s", id), nil)
	if err != nil {
		return
	}

	fmt.Printf("Success add '%s' to your favorite song!!!\n", playingStatus.Item.Name)

	return
}
