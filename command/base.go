// Package command is this application's commands package.
package command

import (
	"fmt"
	"isso0424/spotify_CLI/command/request"
	"isso0424/spotify_CLI/selfmadetypes"
	"isso0424/spotify_CLI/util"
)

var (
	requestCommands = []selfmadetypes.RequestCommand{
		play{},
		pause{},
		status{},
		resume{},
		next{},
		prev{},
		repeat{},
		shuffle{},
		refresh{},
		volume{},
		search{},
		favoriteTrack{},
		importOwnPlaylists{},
		recent{},
		playlist{},
		load{},
		random{},
		addToPlaylist{},
		createPlaylist{},
		deleteTrackFromPlaylist{},
	}

	loadfileCommands = []selfmadetypes.FileloadCommand{
		save{},
		show{},
	}
)

// MainLoop is function that is application's root loop.
func MainLoop(token string) {
	fmt.Println("if you wanna exit, you must type 'exit'")
	err := welcome{}.Execute(&token)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	err = status{}.Execute(&token)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	allCommands := joinCommandList(
		requestCommands,
		loadfileCommands,
	)

	for {
		var commandName string
		util.Input("", "Command", &commandName)

		if commandName == "help" {
			help(allCommands)
		}

		if commandName == "exit" {
			break
		}

		err := execute(&token, commandName)

		if err != nil {
			fmt.Printf("Error: %s", err)
		}
	}
}

func execute(
	token *string,
	commandName string,
) (err error) {
	for _, command := range requestCommands {
		if command.GetCommandName() == commandName {
			err = command.Execute(token)
			if command.GetHelp().Kind == selfmadetypes.Player {
				err = request.PrintPlayingStatus(token)
			}
			return
		}
	}

	for _, command := range loadfileCommands {
		if command.GetCommandName() == commandName {
			err = command.Execute()
			return
		}
	}

	return
}

func joinCommandList(
	requestCommandList []selfmadetypes.RequestCommand,
	loadfileCommandList []selfmadetypes.FileloadCommand,
) (commandList []selfmadetypes.Command) {
	for _, command := range requestCommandList {
		commandList = append(commandList, command)
	}

	for _, command := range loadfileCommandList {
		commandList = append(commandList, command)
	}

	return
}
