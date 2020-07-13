package command

import (
	"fmt"
	"isso0424/spotify_CLI/selfMadeTypes"
	"isso0424/spotify_CLI/util"
)

func load(token string) (err error) {
	fmt.Println("please input playlist name")
	var name string
	util.Input("PlayListName", &name)

	playlistList, err := util.LoadPlayList()

	if err != nil {
		return
	}

	for _, target := range playlistList {
		if target.Name == name {
			fmt.Printf("play %s\n", target.Name)
			play(token, target.Uri)
			return
		}
	}

	err = &selfMadeTypes.NotFound{Target: name}

	return
}
