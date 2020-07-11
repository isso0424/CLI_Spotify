package main

import (
	"fmt"

	"isso0424/spotify-rapspi/command"

	"golang.org/x/oauth2"
)


func main() {
  mainLoop(token)
}

func mainLoop(token *oauth2.Token) {
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
      status := command.GetPlayStatus(token)
      if status {
        fmt.Println("now playing")
      } else {
        fmt.Println("pausing...")
      }
  }
  }
}

