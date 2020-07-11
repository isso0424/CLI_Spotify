package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func Load(token string) {
  if !existFile("playlist.json") {
    fmt.Println("first, you have to save playlist")
    return
  }

  file, err := ioutil.ReadFile("playlist.json")
  if err != nil {
    log.Fatalln("could not read playlist.json")
    return
  }

  var playlistList []playlist

  json.Unmarshal(file, &playlistList)

  fmt.Printf("please input playlist name\nPlayListName|>>")
  var name string
  fmt.Scanln(&name)

  for _, target := range playlistList {
    if target.Name == name {
      fmt.Printf("play %s\n", target.Name)
      play(token, target.Uri)
      return
    }
  }

  fmt.Printf("%s is not found.", name)
}
