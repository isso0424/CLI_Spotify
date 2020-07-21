package command

import "isso0424/spotify_CLI/command/request"

func next(token string) (newToken string, err error) {
	_, newToken, err = request.CreateRequest(token, "POST", "https://api.spotify.com/v1/me/player/next", nil)

	if err != nil {
		return
	}

	_, newToken, err = getPlayStatus(token)

	return
}
