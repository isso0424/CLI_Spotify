package command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"isso0424/spotify-rapspi/selfMadeTypes"
)

func GetPlayStatus(token string) bool {
	response, err := createRequest(token, "GET", "https://api.spotify.com/v1/me/player")
	if err != nil {
		fmt.Println("Error: ", err)
		return false
	}
	if response.StatusCode == 204 {
		fmt.Println("You have to play on spotify client before use this `CLI client`.")
		return false
	}

	buffer := make([]byte, 8192)
	_, err = response.Body.Read(buffer)

	buffer = bytes.Trim(buffer, "\x00")

	var responseBody selfMadeTypes.Content
	if err := json.Unmarshal(buffer, &responseBody); err != nil {
		fmt.Println("Error: ", err)
	}

	createInfo(responseBody)

	return responseBody.IsPlaying && len(responseBody.Item.Artists) != 0
}

func createInfo(content selfMadeTypes.Content) {
	if content.IsPlaying && len(content.Item.Artists) != 0 {
		fmt.Printf("Playing status\n--------------\nTitle: %s\nArtist: %s\n", content.Item.Name, content.Item.Artists[0].Name)
	} else {
		fmt.Println("Pausing")
	}
}
