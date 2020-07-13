package command

import (
	"isso0424/spotify_CLI/auth"
)

func refresh(token string) (string, error) {
	tokenPtr, err := auth.GetToken()
	if err != nil {
		return token, err
	}

	return *tokenPtr, nil
}
