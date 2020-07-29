package request

import "strings"

// GetPlayingPlaylistID is function get playlist id.
func GetPlayingPlaylistID(token *string) (id *string, err error) {
	playingStatus, err := GetStatus(token)
	if err != nil {
		return
	}

	url := playingStatus.Context.Href

	id = &strings.Split(url, "/")[5]

	return
}
