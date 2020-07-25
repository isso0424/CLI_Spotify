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

	for {
		var commandName string
		util.Input("", "Command", &commandName)

		if commandName == "exit" {
			break
		}
		err := execute(&token, commandName, requestCommands)

		if err != nil {
			fmt.Printf("Error: %s", err, requestCommands)
		}
	}
}

func execute(token *string, commandName string, requestCommandList []selfMadeTypes.RequestCommand) (err error) {
  for _, command := range(requestCommandList) {
    if command.GetCommandName() == commandName {
      command.Execute(token)
    }
  }
	switch commandName {
	case "save":
		err = save()
	case "load":
		err = load(token)
	case "show":
		err = show()
	case "refresh":
		err = refresh(token)
	case "random":
		err = random(token)
	}

	return
}
