package command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"isso0424/spotify_CLI/command/parse"
	"isso0424/spotify_CLI/command/request"
	"isso0424/spotify_CLI/selfMadeTypes"
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
