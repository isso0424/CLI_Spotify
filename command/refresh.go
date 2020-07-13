package command

import (
	"fmt"
	"isso0424/spotify_CLI/auth"
)

func refresh() string {
	tokenPtr, err := auth.GetToken()
	if err != nil {
		fmt.Println("Error: ", err)
		return ""
	}

	return *tokenPtr
}
