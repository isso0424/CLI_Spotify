package main

import (
	"fmt"

	"isso0424/spotify-rapspi/auth"
	"isso0424/spotify-rapspi/command"
)


func main() {
  token, err := auth.GetToken()

  if err != nil {
    panic(err)
  }

  mainLoop(token)
}

func mainLoop(token string) {
  fmt.Println("if you wanna exit, you must type 'exit'")
  for {
    var commandKind string
    fmt.Print(">>")
    fmt.Scanln(&commandKind)

    switch commandKind {
    case "exit":
      return
    case "pause":
      command.Pause(token)
      fmt.Println("paused!!!")
    case "resume":
      command.Resume(token)
      fmt.Println("resumed!!!")
    case "status":
      command.GetPlayStatus(token)
  }
  }
}

