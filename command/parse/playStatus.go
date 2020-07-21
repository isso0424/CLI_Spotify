package parse

import (
	"fmt"
	"isso0424/spotify_CLI/selfMadeTypes"
	"strings"
)

func CreatePlayingStatus(content selfMadeTypes.Content, playlist selfMadeTypes.PlayListFromRequest, token string) (status string) {
	if content.IsPlaying && len(content.Item.Artists) != 0 {
		status = fmt.Sprintf("Playing status\n--------------\nTitle: %s\nArtist: %s\n\nPlayList Infomation\n-------------------\nPlayList: %s\nOwner: %s\n", content.Item.Name, content.Item.Artists[0].Name, playlist.Name, playlist.Owner.DisplayName)
	} else {
		status = fmt.Sprintf("Pausing")
	}

	return
}

func getPlaylistID(url string) (*string, error) {
	err := &lengthError{}
	spritted := strings.Split(url, "/")

	if len(spritted) < 5 {
		return nil, err
	}
	tmp := spritted[4]

	id := strings.Split(tmp, "?")[0]

	return &id, nil
}
