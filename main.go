package main

import (
	"fmt"
	"strings"

	"isso0424/spotify_CLI/auth"
	"isso0424/spotify_CLI/command"
	"isso0424/spotify_CLI/util"
)

func main() {
	token, err := auth.GetToken()

	fmt.Println(strings.Split("https://open.spotify.com/playlist/2EKvIPR0K0rvpNnyhKdsd6?si=nUEPESAtSiywv8GkWAHYjg", "/")[3])

	if err != nil {
		panic(err)
	}

	mainLoop(*token)
}

func mainLoop(token string) {
	fmt.Println("if you wanna exit, you must type 'exit'")
	command.GetPlayStatus(token)
	for {
		var commandKind string
		util.Input("Command", &commandKind)

		switch commandKind {
		case "exit":
			return
		case "pause":
			token = command.Pause(token)
		case "resume":
			token = command.Resume(token)
		case "status":
			_, token = command.GetPlayStatus(token)
		case "play":
			token = command.PlayFromURL(token)
		case "save":
			command.Save()
		case "load":
			command.Load(token)
		case "show":
			command.Show()
		case "refresh":
			token = command.Refresh()
		case "random":
			token = command.Random(token)
		case "next":
			token = command.Next(token)
		case "prev":
			token = command.Prev(token)
		case "repeat":
			token = command.Repeat(token)
		}
	}
}
