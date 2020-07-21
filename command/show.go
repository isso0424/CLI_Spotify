package command

import (
	"fmt"
	"isso0424/spotify_CLI/command/file"
)

func show() (err error) {
	playlistList, err := file.LoadPlayList()

	if err != nil {
		return
	}

	for index, target := range playlistList {
		fmt.Printf("id: %d\n------------------------------------------------\nname: %s\nuri: %s\n\n", index, target.Name, target.Uri)
	}

	return
}
