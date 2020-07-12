package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func Show() {
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

  for index, target := range playlistList {
    fmt.Printf("%d\n------------------------------------------------\nname: %s\nuri: %s\n\n", index, target.Name, target.Uri)
  }
}
