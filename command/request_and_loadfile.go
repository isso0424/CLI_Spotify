package command

import (
	"fmt"
	"isso0424/spotify_CLI/command/file"
	"isso0424/spotify_CLI/selfMadeTypes"
	"isso0424/spotify_CLI/util"
)

func(_ random) Execute(token *string) (err error) {
	playlists, err := file.LoadPlayList()
	if err != nil {
		return
	}

	playlist := choice(playlists)

	err = playFromURL(token, playlist.Uri)

	return
}

func(_ load) Execute(token *string) (err error) {
	var name string
	util.Input("please input playlist name", "PlayListName", &name)

	playlistList, err := file.LoadPlayList()

	if err != nil {
		return
	}

	for _, target := range playlistList {
		if target.Name == name {
			fmt.Printf("play %s\n", target.Name)
			err = playFromURL(token, target.Uri)
			return
		}
	}

	err = &selfMadeTypes.NotFound{Target: name}

	return
}
