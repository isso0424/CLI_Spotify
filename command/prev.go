package command

import (
	"fmt"
	"isso0424/spotify_CLI/util"
)

func prev(token string) (newToken string) {
	_, newToken, err := util.CreateRequest(token, "POST", "https://api.spotify.com/v1/me/player/previous", nil)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	_, newToken = getPlayStatus(token)

	return
}
