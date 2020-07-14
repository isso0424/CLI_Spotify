package util

import (
	"encoding/json"
	"io/ioutil"
	"isso0424/spotify_CLI/selfMadeTypes"
	"os"
)

func LoadPlayList() (playlistList []selfMadeTypes.PlayList, err error) {
	if _, err = os.Stat("playlist.json"); err != nil {
		return
	}

	file, err := ioutil.ReadFile("playlist.json")
	if err != nil {
		return
	}

	err = json.Unmarshal(file, &playlistList)

	return
}
