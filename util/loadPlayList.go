package util

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"isso0424/spotify-rapspi/selfMadeTypes"
)

func LoadPlayList() (playlistList []selfMadeTypes.PlayList,err error) {
  if _, err = os.Stat("playlist.json"); err != nil {
    return
  }

  file, err := ioutil.ReadFile("playlist.json")
  if err != nil {
    return
  }

  json.Unmarshal(file, &playlistList)

  return
}
