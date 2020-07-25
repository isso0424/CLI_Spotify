package command

import (
	"fmt"
	"isso0424/spotify_CLI/selfMadeTypes"
	"isso0424/spotify_CLI/util"
)

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

  requestCommands := []selfMadeTypes.RequestCommand{
    play{},
    pause{},
    status{},
    resume{},
    next{},
    prev{},
    repeat{},
    shuffle{},
    refresh{},
  }

  loadfileCommands := []selfMadeTypes.FileloadCommand{
    save{},
    show{},
  }

  requestAndLoadfileCommands := []selfMadeTypes.RequestAndFileloadCommand{
    load{},
    random{},
  }

	for {
		var commandName string
		util.Input("", "Command", &commandName)

		if commandName == "exit" {
			break
		}
		err := execute(&token, commandName, requestCommands, loadfileCommands, requestAndLoadfileCommands)

		if err != nil {
			fmt.Printf("Error: %s", err)
		}
	}
}

func execute(token *string, commandName string, requestCommandList []selfMadeTypes.RequestCommand, loadfileCommandList []selfMadeTypes.FileloadCommand, requestAndLoadfileCommandList []selfMadeTypes.RequestAndFileloadCommand) (err error) {
  for _, command := range(requestCommandList) {
    if command.GetCommandName() == commandName {
      err = command.Execute(token)
      return
    }
  }

  for _, command := range(loadfileCommandList) {
    if command.GetCommandName() == commandName {
      err = command.Execute()
      return
    }
  }

  for _, command := range(requestAndLoadfileCommandList) {
    if command.GetCommandName() == commandName {
      err = command.Execute(token)
      return
    }
  }

	return
}
