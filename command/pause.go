package command

import (
	"fmt"
)

func Pause(token string) (newToken string) {
	_, newToken, err := createRequest(token, "PUT", "https://api.spotify.com/v1/me/player/pause", nil)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("paused!!!")

	return
}
