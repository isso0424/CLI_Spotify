package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"isso0424/spotify-rapspi/util"
)

func Load(token string) {
  if !existFile("playlist.json") {
    fmt.Println("Error: first, you have to save playlist")
    return
  }

  file, err := ioutil.ReadFile("playlist.json")
  if err != nil {
    fmt.Println("Error: could not read playlist.json")
    return
  }

  var playlistList []playlist

  json.Unmarshal(file, &playlistList)

  fmt.Println("please input playlist name")
  var name string
  util.Input("PlayListName", &name)

  for _, target := range playlistList {
    if target.Name == name {
      fmt.Printf("play %s\n", target.Name)
      play(token, target.Uri)
      return
    }
  }

  fmt.Printf("Error: %s is not found.", name)
}
