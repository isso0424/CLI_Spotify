package command

import (
	"fmt"
	"isso0424/spotify_CLI/util"
)

func resume(token string) (newToken string, err error) {
	_, newToken, err = util.CreateRequest(token, "PUT", "https://api.spotify.com/v1/me/player/play", nil)

	if err != nil {
		return
	}
	fmt.Println("resumed!!!")

	return
}
