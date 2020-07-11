package main

import (
	"fmt"
	"strings"

	"isso0424/spotify-rapspi/auth"
	"isso0424/spotify-rapspi/command"
	"isso0424/spotify-rapspi/util"
)


func main() {
  token, err := auth.GetToken(false)

  fmt.Println(strings.Split("https://open.spotify.com/playlist/2EKvIPR0K0rvpNnyhKdsd6?si=nUEPESAtSiywv8GkWAHYjg", "/")[3])

  if err != nil {
    panic(err)
  }

  mainLoop(token)
}

func mainLoop(token string) {
  fmt.Println("if you wanna exit, you must type 'exit'")
  command.GetPlayStatus(token)
  for {
    var commandKind string
    util.Input("Command", &commandKind)

    switch commandKind {
    case "exit":
      return
    case "pause":
      command.Pause(token)
    case "resume":
      command.Resume(token)
    case "status":
      command.GetPlayStatus(token)
    case "play":
      command.PlayFromURL(token)
    case "save":
      command.Save()
    case "load":
      command.Load(token)
    case "show":
      command.Show()
    case "refresh":
      token, _ = auth.GetToken(true)
  }
  }
}

