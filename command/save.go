package command

import (
	"fmt"
	"isso0424/spotify_CLI/command/file"
	"isso0424/spotify_CLI/command/parse"
	"isso0424/spotify_CLI/selfMadeTypes"
	"isso0424/spotify_CLI/util"
)

func save() (err error) {
	fmt.Printf("please input playlist url\n")
	var url string
	util.Input("PlayListURL", &url)

	uri, err := parse.CreateContextUri(url)
	if err != nil {
		return
	}

	fmt.Printf("\nplease input playlist name\n")
	var name string
	util.Input("PlayListName", &name)

	list := selfMadeTypes.PlayList{Uri: *uri, Name: name}

	if checkDuplicateName(name) {
		err = file.SavePlayList(list)
	} else {
    err = &selfMadeTypes.NameDuplicateError{Target: name}
	}

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
