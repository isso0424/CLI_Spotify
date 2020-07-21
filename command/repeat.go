package command

import (
	"fmt"
	"isso0424/spotify_CLI/command/request"
)

func repeat(token string) (newToken string, err error) {
	newToken = token

	status, newToken, err := getStatus(token)

	if err != nil {
		return
	}

	state := switchRepeatState(status.RepeatState)

	_, newToken, err = request.CreateRequest(token, "PUT", fmt.Sprintf("https://api.spotify.com/v1/me/player/repeat?state=%s", state), nil)

	if err != nil {
    return
	}

	fmt.Printf("Repeat state change to `%s`\n", state)

	return
}

func switchRepeatState(state string) string {
	switch state {
	case "track":
		return "off"
	case "context":
		return "track"
	case "off":
		return "context"
	}

	return "off"
}
