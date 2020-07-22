package command

import (
	"isso0424/spotify_CLI/command/file"
	"isso0424/spotify_CLI/command/parse"
	"isso0424/spotify_CLI/selfMadeTypes"
	"isso0424/spotify_CLI/util"
)

func save() (err error) {
	var url string
	util.Input("please input playlist url\n", "PlayListURL", &url)

	uri, err := parse.CreateContextUri(url)
	if err != nil {
		return
	}

	var name string
	util.Input("\nplease input playlist name\n", "PlayListName", &name)

	list := selfMadeTypes.PlayList{Uri: *uri, Name: name}

	playlistList, err := file.LoadPlayList()

  if err != nil {
    return
  }

	if checkDuplicateName(name, playlistList) {
		err = file.SavePlayList(list)
	} else {
		err = &selfMadeTypes.NameDuplicateError{Target: name}
	}

	return
}

func checkDuplicateName(name string, playlistList []selfMadeTypes.PlayList) bool {
	for _, content := range playlistList {
		if content.Name == name {
			return false
		}
	}

	return true
}
