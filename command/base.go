package command

import (
	"fmt"
	"isso0424/spotify_CLI/util"
)


func MainLoop(token string) {
	fmt.Println("if you wanna exit, you must type 'exit'")
	getPlayStatus(token)
	for {
		var commandName string
		util.Input("Command", &commandName)

    if commandName == "exit" {
      break
    }
    token = command(token, commandName)
	}
}

func command(token string, commandName string) (newToken string) {
  newToken = token
	switch commandName {
	case "pause":
		newToken = pause(newToken)
	case "resume":
		newToken = resume(newToken)
	case "status":
		_, newToken = getPlayStatus(newToken)
	case "play":
		newToken = playFromURL(newToken)
	case "save":
		save()
	case "load":
		load(newToken)
	case "show":
		show()
	case "refresh":
		newToken = refresh()
	case "random":
		newToken = random(newToken)
	case "next":
		newToken = next(newToken)
	case "prev":
		newToken = prev(newToken)
	case "repeat":
		newToken = repeat(newToken)
	case "shuffle":
		newToken = shuffle(newToken)
	}

  return
}
