package command

import (
	"fmt"
	"isso0424/spotify-rapspi/util"
)

func Load(token string) {
  fmt.Println("please input playlist name")
  var name string
  util.Input("PlayListName", &name)

  playlistList, _ := util.LoadPlayList()

  for _, target := range playlistList {
    if target.Name == name {
      fmt.Printf("play %s\n", target.Name)
      play(token, target.Uri)
      return
    }
  }

  fmt.Printf("Error: %s is not found.", name)
}
