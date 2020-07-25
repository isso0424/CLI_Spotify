package command

import (
	"isso0424/spotify_CLI/auth"
)

func(_ refresh) Execute(token *string) error {
	tokenPtr, err := auth.GetToken()
	if err != nil {
		return err
	}

	*token = *tokenPtr

	return nil
}
