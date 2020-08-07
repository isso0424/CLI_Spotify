// Package command is this application's commands package.
package command

import (
	"fmt"
	"isso0424/spotify_CLI/command/request"
	"isso0424/spotify_CLI/selfmadetypes/commandtypes"
	"isso0424/spotify_CLI/util"
)

var (
	requestCommands = []commandtypes.RequestCommand{
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

	loadfileCommands = []commandtypes.FileloadCommand{
		save{},
		show{},
	}
)

// MainLoop is function that is application's root loop.
func MainLoop(token string) {
	fmt.Println("If you wanna exit, you must execute 'exit'")
	fmt.Println("If you wanna get commands help, you must execute 'help'")
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
		commandName := util.Input("", "Command")

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
			if command.GetHelp().Kind == commandtypes.Player {
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
	requestCommandList []commandtypes.RequestCommand,
	loadfileCommandList []commandtypes.FileloadCommand,
) (commandList []commandtypes.Command) {
	for _, command := range requestCommandList {
		commandList = append(commandList, command)
	}

	for _, command := range loadfileCommandList {
		commandList = append(commandList, command)
	}

	return
}
