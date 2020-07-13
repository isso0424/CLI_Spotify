package main

import (
	"isso0424/spotify_CLI/auth"
	"isso0424/spotify_CLI/command"
)

func main() {
	token, err := auth.GetToken()

	if err != nil {
		panic(err)
	}

	command.MainLoop(*token)
}

