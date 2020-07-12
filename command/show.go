package command

import (
	"fmt"
	"isso0424/spotify_CLI/util"
)

func Show() {
	playlistList, _ := util.LoadPlayList()

	for index, target := range playlistList {
		fmt.Printf("%d\n------------------------------------------------\nname: %s\nuri: %s\n\n", index, target.Name, target.Uri)
	}
}
