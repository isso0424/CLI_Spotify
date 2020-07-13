package command

import (
	"fmt"
	"isso0424/spotify_CLI/util"
)


func MainLoop(token string) {
  var err error
	fmt.Println("if you wanna exit, you must type 'exit'")
	getPlayStatus(token)
	for {
		var commandName string
		util.Input("Command", &commandName)

    if commandName == "exit" {
      break
    }
    token, err = command(token, commandName)

    if err != nil {
      fmt.Printf("Error: %s", err)
      err = nil
    }
	}
}

func command(token string, commandName string) (newToken string, err error) {
  newToken = token
	switch commandName {
	case "pause":
		newToken, err = pause(newToken)
	case "resume":
		newToken, err = resume(newToken)
	case "status":
		_, newToken, err = getPlayStatus(newToken)
	case "play":
		newToken, err = playFromURL(newToken)
	case "save":
		err = save()
	case "load":
		err = load(newToken)
	case "show":
		err = show()
	case "refresh":
		newToken, err = refresh(token)
	case "random":
		newToken, err = random(newToken)
	case "next":
		newToken, err = next(newToken)
	case "prev":
		newToken, err = prev(newToken)
	case "repeat":
		newToken, err = repeat(newToken)
	case "shuffle":
		newToken, err = shuffle(newToken)
	}

  return
}
