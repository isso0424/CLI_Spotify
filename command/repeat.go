package command

import (
	"fmt"
	"isso0424/spotify_CLI/util"
)

func repeat(token string) (newToken string) {
	newToken = token

	status, newToken, err := getStatus(token)

	if status == nil {
		fmt.Println("Error: Failed to get playing status")
		return
	}

	state := switchRepeatState(status.RepeatState)

	_, newToken, err = util.CreateRequest(token, "PUT", fmt.Sprintf("https://api.spotify.com/v1/me/player/repeat?state=%s", state), nil)

	if err != nil {
		fmt.Println("Error: ", err)
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
