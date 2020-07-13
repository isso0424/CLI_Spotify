package command

import (
	"fmt"
	"isso0424/spotify_CLI/util"
)

func Next(token string) (newToken string) {
	_, newToken, err := util.CreateRequest(token, "POST", "https://api.spotify.com/v1/me/player/next", nil)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	GetPlayStatus(token)

	return
}
