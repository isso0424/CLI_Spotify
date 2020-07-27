// Package command is this application's commands package.
package command

import (
	"fmt"
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
	}

	loadfileCommands = []selfmadetypes.FileloadCommand{
		save{},
		show{},
	}

	requestAndLoadfileCommands = []selfmadetypes.RequestAndFileloadCommand{
		load{},
		random{},
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
		requestAndLoadfileCommands,
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
			return
		}
	}

	for _, command := range loadfileCommands {
		if command.GetCommandName() == commandName {
			err = command.Execute()
			return
		}
	}

	for _, command := range requestAndLoadfileCommands {
		if command.GetCommandName() == commandName {
			err = command.Execute(token)
			return
		}
	}

	return
}

func joinCommandList(
	requestCommandList []selfmadetypes.RequestCommand,
	loadfileCommandList []selfmadetypes.FileloadCommand,
	requestAndLoadfileCommandList []selfmadetypes.RequestAndFileloadCommand,
) (commandList []selfmadetypes.Command) {
	for _, command := range requestCommandList {
		commandList = append(commandList, command)
	}

	for _, command := range loadfileCommandList {
		commandList = append(commandList, command)
	}

	for _, command := range requestAndLoadfileCommandList {
		commandList = append(commandList, command)
	}

	return
}
