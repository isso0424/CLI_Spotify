package parse

import (
	"fmt"
	"isso0424/spotify_CLI/selfMadeTypes"
)

func CreatePlayingStatus(content selfMadeTypes.Content, playlist selfMadeTypes.PlayListFromRequest) (status string) {
	if content.IsPlaying && len(content.Item.Artists) != 0 {
		status = fmt.Sprintf(
			"Playing status\n"+
				"--------------\n"+
				"Title: %s\n"+
				"Artist: %s\n\n"+
				"PlayList Infomation\n"+
				"-------------------\n"+
				"PlayList: %s\n"+
				"Owner: %s\n",
			content.Item.Name,
			content.Item.Artists[0].Name,
			playlist.Name,
			playlist.Owner.DisplayName,
		)
	} else {
		status = "Pausing"
	}

	return
}
