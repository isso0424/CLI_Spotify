package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"isso0424/spotify_CLI/selfMadeTypes"
	"isso0424/spotify_CLI/util"
)

func save() (err error) {
	fmt.Printf("please input playlist url\n")
	var url string
	util.Input("PlayListURL", &url)

	uri, err := CreateContextUri(url)
	if err != nil {
		return
	}

	fmt.Printf("\nplease input playlist name\n")
	var name string
	util.Input("PlayListName", &name)

	list := selfMadeTypes.PlayList{Uri: *uri, Name: name}

	if checkDuplicateName(name) {
		err = saveToJson(list)
	} else {
    err = &selfMadeTypes.NameDuplicateError{name}
	}

  return
}

func saveToJson(target selfMadeTypes.PlayList) (err error) {
	playlistList, _ := util.LoadPlayList()
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

func checkDuplicateName(name string) bool {
	playlistList, err := util.LoadPlayList()
	if err != nil {
		return true
	}

	for _, content := range playlistList {
		if content.Name == name {
			return false
		}
	}

	return true
}
