package command

import (
	"fmt"
	"isso0424/spotify_CLI/command/file"
	"isso0424/spotify_CLI/selfmadetypes"
	"isso0424/spotify_CLI/util"
)

// Execute is excution command function.
func (cmd random) Execute(token *string) (err error) {
	playlists, err := file.LoadPlayList()
	if err != nil {
		return
	}

	choisePlaylist := util.Choose(playlists)
	err = playFromURL(token, choisePlaylist.URI)

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
			err = playFromURL(token, target.URI)
			return
		}
	}

	err = &selfmadetypes.NotFound{Target: name}

	return
}
