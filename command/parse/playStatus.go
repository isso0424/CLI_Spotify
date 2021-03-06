// Package parse is responseTypes parser.
package parse

import (
	"fmt"
	"isso0424/spotify_CLI/selfmadetypes"
	"isso0424/spotify_CLI/selfmadetypes/responsetypes"
)

// CreatePlayingStatus is parsing play status function
func CreatePlayingStatus(
	content responsetypes.Content,
	contextName string,
	contextUser string,
	kind string,
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
					"Playing " + kind,
				},
				getContextStatus(contextUser, contextName, kind),
				{
					fmt.Sprintf("Shuffle: %v", content.ShuffleState),
					"Repeat: " + content.RepeatState,
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

func getContextStatus(contextUser, contextName, kind string) (status []string) {
	switch kind {
	case "playlist":
		status = []string{
			"Playlist name: " + contextName,
			"Owner: " + contextUser,
		}
	case "artist":
		status = []string{
			"Artist name: " + contextName,
		}
	case "album":
		status = []string{
			"Album name: " + contextName,
			"Artist name: " + contextUser,
		}
	default:
		status = []string{}
	}

	return
}
