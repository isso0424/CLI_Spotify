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
  }

  loadfileCommands := []selfMadeTypes.FileloadCommand{
    save{},
    show{},
  }

	for {
		var commandName string
		util.Input("", "Command", &commandName)

		if commandName == "exit" {
			break
		}
		err := execute(&token, commandName, requestCommands, loadfileCommands)

		if err != nil {
			fmt.Printf("Error: %s", err, requestCommands)
		}
	}
}

func execute(token *string, commandName string, requestCommandList []selfMadeTypes.RequestCommand, loadfileCommandList []selfMadeTypes.FileloadCommand) (err error) {
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
	switch commandName {
	case "load":
		err = load(token)
	case "refresh":
		err = refresh(token)
	case "random":
		err = random(token)
	}

	return
}
