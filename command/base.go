package command

import (
	"fmt"
	"isso0424/spotify_CLI/util"
)

func MainLoop(token string) {
	fmt.Println("if you wanna exit, you must type 'exit'")
	err := welcome(&token)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	_, err = getPlayStatus(&token)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	for {
		var commandName string
		util.Input("", "Command", &commandName)

		if commandName == "exit" {
			break
		}
		err := command(&token, commandName)

		if err != nil {
			fmt.Printf("Error: %s", err)
		}
	}
}

func command(token *string, commandName string) (err error) {
	switch commandName {
	case "pause":
		err = pause(token)
	case "resume":
		err = resume(token)
	case "status":
		_, err = getPlayStatus(token)
	case "play":
		err = playFromURL(token)
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
	case "next":
		err = next(token)
	case "prev":
		err = prev(token)
	case "repeat":
		err = repeat(token)
	case "shuffle":
		err = shuffle(token)
	}

	return
}
