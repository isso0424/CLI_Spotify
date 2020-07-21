package command

import (
	"fmt"
	"isso0424/spotify_CLI/util"
)

func show() (err error) {
	playlistList, err := util.LoadPlayList()

	if err != nil {
		return
	}

	for index, target := range playlistList {
    fmt.Printf("id: %d\n------------------------------------------------\nname: %s\nuri: %s\n\n", index, target.Name, target.Uri)
	}

	return
}
