package command

import (
	"fmt"
)

func Resume(token string) (newToken string) {
	_, newToken, err := createRequest(token, "PUT", "https://api.spotify.com/v1/me/player/play", nil)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("resumed!!!")

	return
}
