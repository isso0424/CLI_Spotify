// Package parse is response parser.
package parse

import (
	"isso0424/spotify_CLI/selfmadetypes"
)

// CreatePlayingStatus is parsing play status function
func CreatePlayingStatus(
	content selfmadetypes.Content,
	playlist selfmadetypes.PlayList,
) (status selfmadetypes.OutputMessage) {
	if content.IsPlaying && len(content.Item.Artists) != 0 {
		status = selfmadetypes.OutputMessage{
			Message: [][]string{
				{
					"Playing status",
				},
				{
					"Title: " + content.Item.Name,
					"Artist: " + content.Item.Artists[0].Name,
				},
				{
					"Playing Item",
				},
				{
					"SearchResultItem: " + playlist.Name,
					"Owner: " + playlist.Owner.DisplayName,
				},
			},
		}
	} else {
		status = selfmadetypes.OutputMessage{
			Message: [][]string{
				{
					"Pausing",
				},
			},
		}
	}

	return status
}
