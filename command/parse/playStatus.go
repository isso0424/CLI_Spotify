package parse

import (
	"fmt"
	"isso0424/spotify_CLI/selfMadeTypes"
)

func CreatePlayingStatus(content selfMadeTypes.Content, playlist selfMadeTypes.PlayListFromRequest, token string) (status string) {
	if content.IsPlaying && len(content.Item.Artists) != 0 {
		status = fmt.Sprintf("Playing status\n--------------\nTitle: %s\nArtist: %s\n\nPlayList Infomation\n-------------------\nPlayList: %s\nOwner: %s\n", content.Item.Name, content.Item.Artists[0].Name, playlist.Name, playlist.Owner.DisplayName)
	} else {
		status = "Pausing"
	}

	return
}
