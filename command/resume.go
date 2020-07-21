package command

import (
	"fmt"
	"isso0424/spotify_CLI/command/request"
)

func resume(token string) (newToken string, err error) {
	_, newToken, err = request.CreateRequest(token, "PUT", "https://api.spotify.com/v1/me/player/play", nil)

	if err != nil {
		return
	}
	fmt.Println("resumed!!!")

	return
}
