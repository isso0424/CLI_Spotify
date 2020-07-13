package command

import (
	"fmt"
	"isso0424/spotify_CLI/util"
)


func MainLoop(token string) {
	fmt.Println("if you wanna exit, you must type 'exit'")
	GetPlayStatus(token)
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
		newToken = Pause(newToken)
	case "resume":
		newToken = Resume(newToken)
	case "status":
		_, newToken = GetPlayStatus(newToken)
	case "play":
		newToken = PlayFromURL(newToken)
	case "save":
		Save()
	case "load":
		Load(newToken)
	case "show":
		Show()
	case "refresh":
		newToken = Refresh()
	case "random":
		newToken = Random(newToken)
	case "next":
		newToken = Next(newToken)
	case "prev":
		newToken = Prev(newToken)
	case "repeat":
		newToken = Repeat(newToken)
	case "shuffle":
		newToken = Shuffle(newToken)
	}

  return
}
