package command

import (
	"isso0424/spotify_CLI/util"
)

func prev(token string) (newToken string, err error) {
	_, newToken, err = util.CreateRequest(token, "POST", "https://api.spotify.com/v1/me/player/previous", nil)

	if err != nil {
		return
	}

	_, newToken, err = getPlayStatus(token)

	return
}
