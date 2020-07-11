package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func Save() {
  fmt.Printf("please input playlist url\nPlayListURL|>>>")
  var url string
  fmt.Scanln(&url)
  uri, err := CreateContextUri(url)
  if err != nil {
    log.Fatalln(err)
    return
  }

  fmt.Printf("\nplease input playlist name\nPlayListName|>>>")
  var name string
  fmt.Scanln(&name)

  list := playlist{Uri: *uri, Name: name}

  saveToJson(list)
}

func saveToJson(target playlist) {
  var playlistList []playlist
  if existFile("playlist.json") {
    file, err := ioutil.ReadFile("playlist.json")
    if err != nil {
      log.Fatalln("could not read playlist.json")
      return
    }

    json.Unmarshal(file, &playlistList)
  }
  playlistList = append(playlistList, target)

  jsonFile, err := json.Marshal(playlistList)
  if err != nil {
    log.Fatalln(err)
    return
  }

  err = ioutil.WriteFile("playlist.json", jsonFile, 0666)

  if err != nil {
    log.Fatalln(err)
  }

  fmt.Printf("\nplaylist saved!!!\nurl: %s\nname: %s\n", target.Uri, target.Name)
}

func existFile(fileName string) bool {
  _, err := os.Stat(fileName)
  return err == nil
}

type playlist struct {
  Uri string `json:"uri"`
  Name string `json:"name"`
}
