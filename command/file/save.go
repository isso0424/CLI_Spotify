package file

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"isso0424/spotify_CLI/selfMadeTypes"
)

func SavePlayList(target selfMadeTypes.PlayList) (err error) {
	playlistList, _ := LoadPlayList()
	playlistList = append(playlistList, target)

	jsonFile, err := json.Marshal(playlistList)
	if err != nil {
		return
	}

	err = ioutil.WriteFile("playlist.json", jsonFile, 0666)

	if err != nil {
		return
	}

	fmt.Printf("\nplaylist saved!!!\nurl: %s\nname: %s\n", target.Uri, target.Name)

	return
}
