package command

import (
	"fmt"
	"isso0424/spotify_CLI/selfmadetypes"
	"isso0424/spotify_CLI/util"
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

	requestCommands := []selfmadetypes.RequestCommand{
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

	loadfileCommands := []selfmadetypes.FileloadCommand{
		save{},
		show{},
	}

	requestAndLoadfileCommands := []selfmadetypes.RequestAndFileloadCommand{
		load{},
		random{},
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

		err := execute(&token, commandName, requestCommands, loadfileCommands, requestAndLoadfileCommands)

		if err != nil {
			fmt.Printf("Error: %s", err)
		}
	}
}

func execute(
	token *string,
	commandName string,
	requestCommandList []selfmadetypes.RequestCommand,
	loadfileCommandList []selfmadetypes.FileloadCommand,
	requestAndLoadfileCommandList []selfmadetypes.RequestAndFileloadCommand,
) (err error) {
	for _, command := range requestCommandList {
		if command.GetCommandName() == commandName {
			err = command.Execute(token)
			return
		}
	}

	for _, command := range loadfileCommandList {
		if command.GetCommandName() == commandName {
			err = command.Execute()
			return
		}
	}

	for _, command := range requestAndLoadfileCommandList {
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
