package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"isso0424/spotify_CLI/selfMadeTypes"
	"isso0424/spotify_CLI/util"
)

func save() {
	fmt.Printf("please input playlist url\n")
	var url string
	util.Input("PlayListURL", &url)
	uri, err := CreateContextUri(url)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Printf("\nplease input playlist name\n")
	var name string
	util.Input("PlayListName", &name)

	list := selfMadeTypes.PlayList{Uri: *uri, Name: name}

	if checkDuplicateName(name) {
		saveToJson(list)
	} else {
		fmt.Println("Error: This name is duplicated.")
	}
}

func saveToJson(target selfMadeTypes.PlayList) {
	playlistList, _ := util.LoadPlayList()
	playlistList = append(playlistList, target)

	jsonFile, err := json.Marshal(playlistList)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	err = ioutil.WriteFile("playlist.json", jsonFile, 0666)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Printf("\nplaylist saved!!!\nurl: %s\nname: %s\n", target.Uri, target.Name)
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
