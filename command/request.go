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

func getPlayStatus(token string) (bool, string, error) {
	status, newToken, err := getStatus(token)
	if err != nil {
		return false, newToken, err
	}

	if status == nil {
		return false, newToken, nil
	}
	playlistUrl := status.Context.ExternalUrls.Spotify
	playlistID, err := parse.GetPlaylistID(playlistUrl)

	if err != nil {
		return false, newToken, err
	}

	playListStatus, err := request.GetPlayListStatus(newToken, playlistID)

	if err != nil {
		return false, newToken, err
	}

	fmt.Println(parse.CreatePlayingStatus(*status, playListStatus))

	return status.IsPlaying && len(status.Item.Artists) != 0, newToken, nil
}

func getStatus(token string) (status *selfMadeTypes.Content, newToken string, err error) {
	response, newToken, err := request.CreateRequest(token, "GET", "https://api.spotify.com/v1/me/player", nil)
	if err != nil {
		return
	}
	if response.StatusCode == 204 {
		err = &selfMadeTypes.FailedGetError{Target: "playing status"}
		return
	}

	buffer := make([]byte, 8192)
	_, err = response.Body.Read(buffer)
	if err != nil {
		return
	}

	buffer = bytes.Trim(buffer, "\x00")

	err = json.Unmarshal(buffer, &status)

	return
}

func next(token string) (newToken string, err error) {
	_, newToken, err = request.CreateRequest(token, "POST", "https://api.spotify.com/v1/me/player/next", nil)

	if err != nil {
		return
	}

	_, newToken, err = getPlayStatus(token)

	return
}

func pause(token string) (newToken string, err error) {
	_, newToken, err = request.CreateRequest(token, "PUT", "https://api.spotify.com/v1/me/player/pause", nil)

	if err != nil {
		return
	}
	fmt.Println("paused!!!")

	return
}

func playFromURL(token string) (newToken string, err error) {
	newToken = token

	var url string
	util.Input("please input playlist url\n------------------------", "PlayListURL", &url)

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

func prev(token string) (newToken string, err error) {
	_, newToken, err = request.CreateRequest(token, "POST", "https://api.spotify.com/v1/me/player/previous", nil)

	if err != nil {
		return
	}

	_, newToken, err = getPlayStatus(token)

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
func repeat(token string) (newToken string, err error) {
	newToken = token

	status, newToken, err := getStatus(token)

	if err != nil {
		return
	}

	state := switchRepeatState(status.RepeatState)

	_, newToken, err = request.CreateRequest(token, "PUT", fmt.Sprintf("https://api.spotify.com/v1/me/player/repeat?state=%s", state), nil)

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

func resume(token string) (newToken string, err error) {
	_, newToken, err = request.CreateRequest(token, "PUT", "https://api.spotify.com/v1/me/player/play", nil)

	if err != nil {
		return
	}
	fmt.Println("resumed!!!")

	return
}

func shuffle(token string) (newToken string, err error) {
	newToken = token

	status, newToken, err := getStatus(token)
	if err != nil {
		return
	}

	state := !status.ShuffleState

	_, newToken, err = request.CreateRequest(token, "PUT", fmt.Sprintf("https://api.spotify.com/v1/me/player/shuffle?state=%v", state), nil)
	if err != nil {
		return
	}

	fmt.Printf("Switch shuffle state to %v\n", state)

	return
}
