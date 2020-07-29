package command

import (
	"fmt"
	"isso0424/spotify_CLI/command/file"
	"isso0424/spotify_CLI/command/parse"
	"isso0424/spotify_CLI/command/request"
	"isso0424/spotify_CLI/selfmadetypes"
	"isso0424/spotify_CLI/util"
)

// Execute is excution command function.
func (cmd save) Execute() (err error) {
	var url string
	util.Input("please input playlist url\n", "PlayListURL", &url)

	uri, err := parse.CreateContextURI(url)
	if err != nil {
		return
	}

	var name string
	util.Input("\nplease input playlist name\n", "PlayListName", &name)

	list := selfmadetypes.PlayList{URI: *uri, Name: name}

	playlistList, err := file.LoadPlayList()

	if err != nil {
		return
	}

	if util.CheckDuplicateName(name, playlistList) {
		err = file.SavePlayList(list)
	} else {
		err = &selfmadetypes.NameDuplicateError{Target: name}
	}

	return
}

// Execute is excution command function.
func (cmd show) Execute() (err error) {
	playlistList, err := file.LoadPlayList()

	if err != nil {
		return
	}

	for index, target := range playlistList {
		fmt.Printf(
			"id: %d\n------------------------------------------------\nname: %s\nuri: %s\n\n",
			index,
			target.Name,
			target.URI,
		)
	}

	return
}

// Execute is excution command function.
func (cmd random) Execute(token *string) (err error) {
	playlists, err := file.LoadPlayList()
	if err != nil {
		return
	}

	choisePlaylist := util.Choose(playlists)
	err = request.PlayFromURL(token, choisePlaylist.URI)

	return
}

// Execute is excution command function.
func (cmd load) Execute(token *string) (err error) {
	var name string
	util.Input("please input playlist name", "PlayListName", &name)

	playlistList, err := file.LoadPlayList()

	if err != nil {
		return
	}

	for _, target := range playlistList {
		if target.Name == name {
			fmt.Printf("play %s\n", target.Name)
			err = request.PlayFromURL(token, target.URI)
			return
		}
	}

	err = &selfmadetypes.NotFound{Target: name}

	return
}
