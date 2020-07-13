package command

import (
	"isso0424/spotify_CLI/util"
)

func next(token string) (newToken string, err error) {
	_, newToken, err = util.CreateRequest(token, "POST", "https://api.spotify.com/v1/me/player/next", nil)

	if err != nil {
		return
	}

	getPlayStatus(token)

	return
}
