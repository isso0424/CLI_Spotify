package command

import (
	"fmt"
	"isso0424/spotify_CLI/util"
)

func shuffle(token string) (newToken string, err error) {
	newToken = token

	status, newToken, err := getStatus(token)

	if err != nil {
		return
	}

	state := switchShuffleState(status.ShuffleState)

	_, newToken, err = util.CreateRequest(token, "PUT", fmt.Sprintf("https://api.spotify.com/v1/me/player/shuffle?state=%v", state), nil)

	if err != nil {
		return
	}

	fmt.Printf("Switch shuffle state to %v\n", state)

	return
}

func switchShuffleState(prevState bool) bool {
	return !prevState
}
