package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"isso0424/spotify-rapspi/util"
	"log"
	"os"
)

func Save() {
  fmt.Printf("please input playlist url\n")
  var url string
  util.Input("PlayListURL", &url)
  uri, err := CreateContextUri(url)
  if err != nil {
    log.Fatalln(err)
    return
  }

  fmt.Printf("\nplease input playlist name\n")
  var name string
  util.Input("PlayListName", &name)

  list := playlist{Uri: *uri, Name: name}

  if checkDuplicateName(name) {
    saveToJson(list)
  } else {
    fmt.Println("This name is duplicated.")
  }
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

func checkDuplicateName(name string) bool {
  if !existFile("playlist.json") {
    return true
  }
  file, err := ioutil.ReadFile("playlist.json")
  if err != nil {
    log.Fatalln("could not read playlist.json")
    return false
  }

  var playlistList []playlist

  json.Unmarshal(file, &playlistList)

  for _, content := range playlistList {
    if content.Name == name {
      return false
    }
  }

  return true
}

func existFile(fileName string) bool {
  _, err := os.Stat(fileName)
  return err == nil
}

type playlist struct {
  Uri string `json:"uri"`
  Name string `json:"name"`
}
