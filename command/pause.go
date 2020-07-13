package command

import (
	"fmt"
	"isso0424/spotify_CLI/util"
)

func Pause(token string) (newToken string) {
	_, newToken, err := util.CreateRequest(token, "PUT", "https://api.spotify.com/v1/me/player/pause", nil)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("paused!!!")

	return
}
