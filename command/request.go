package command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"isso0424/spotify_CLI/command/parse"
	"isso0424/spotify_CLI/command/request"
	"isso0424/spotify_CLI/selfMadeTypes"
	"isso0424/spotify_CLI/util"
	"math/rand"
	"time"
)

func(_ status) execute(token *string) (error) {
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


func(_ next) execute(token *string) (err error) {
	_, err = request.CreateRequest(token, selfMadeTypes.POST, "/me/player/next", nil)

	if err != nil {
		return
	}

	err = status{}.execute(token)

	return
}

func(_ pause) execute(token *string) (err error) {
	_, err = request.CreateRequest(token, selfMadeTypes.PUT, "/me/player/pause", nil)

	if err != nil {
		return
	}
	fmt.Println("paused!!!")

	return
}

func(_ play) execute(token *string) (err error) {
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
	values, err := json.Marshal(playJson{ContextUri: uri})
	if err != nil {
		return
	}

	_, err = request.CreateRequest(token, selfMadeTypes.PUT, "/me/player/play", bytes.NewBuffer(values))

	if err != nil {
		return
	}

	err = status{}.execute(token)

	if err != nil {
		return
	}

	return
}

func(_ prev) execute(token *string) (err error) {
	_, err = request.CreateRequest(token, selfMadeTypes.POST, "/me/player/previous", nil)

	if err != nil {
		return
	}

	err = status{}.execute(token)

	return
}

type playJson struct {
	ContextUri string `json:"context_uri"`
}

func choice(playlists []selfMadeTypes.PlayList) selfMadeTypes.PlayList {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(playlists))

	return playlists[index]
}
func(_ repeat) execute(token *string) (err error) {
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

func(_ resume) execute(token *string) (err error) {
	_, err = request.CreateRequest(token, selfMadeTypes.PUT, "/me/player/play", nil)

	if err != nil {
		return
	}
	fmt.Println("resumed!!!")

	return
}

func(_ shuffle) execute(token *string) (err error) {
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

func(_ welcome) execute(token *string) (err error) {
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
