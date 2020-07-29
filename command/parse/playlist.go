package parse

import (
	"isso0424/spotify_CLI/command/request"
	"strings"
)

func GetPlayingPlaylistID(token *string) (id *string, err error) {
	playingStatus, err := request.GetStatus(token)
	if err != nil {
		return
	}

	url := playingStatus.Context.Href

	id = &strings.Split(url, "/")[5]

	return
}
