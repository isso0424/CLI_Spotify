package command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"isso0424/spotify_CLI/selfMadeTypes"
	"isso0424/spotify_CLI/util"
)

func GetPlayStatus(token string) (bool, string) {
	status, newToken, err := getStatus(token)
	if err != nil {
		fmt.Println("Error: ", err)
		return false, newToken
	}

	if status == nil {
		return false, newToken
	}

	createInfo(*status)

	return status.IsPlaying && len(status.Item.Artists) != 0, newToken
}

func getStatus(token string) (status *selfMadeTypes.Content, newToken string, err error) {
	response, newToken, err := util.CreateRequest(token, "GET", "https://api.spotify.com/v1/me/player", nil)
	if err != nil {
		return
	}
	if response.StatusCode == 204 {
		fmt.Println("You have to play on spotify client before use this `CLI client`.")
		return
	}

	buffer := make([]byte, 8192)
	_, err = response.Body.Read(buffer)

	buffer = bytes.Trim(buffer, "\x00")

	if err := json.Unmarshal(buffer, &status); err != nil {
		fmt.Println("Error: ", err)
	}

	return
}

func createInfo(content selfMadeTypes.Content) {
	if content.IsPlaying && len(content.Item.Artists) != 0 {
		fmt.Printf("Playing status\n--------------\nTitle: %s\nArtist: %s\n", content.Item.Name, content.Item.Artists[0].Name)
	} else {
		fmt.Println("Pausing")
	}
}
