package parse

import (
	"isso0424/spotify_CLI/selfmadetypes/responseTypes"
)

// GetArtistNames is extract artist names from artists object slice.
func GetArtistNames(artists []responseTypes.Artists) (artistNames string) {
	artistNames = ""

	for _, artist := range artists {
		artistNames += artist.Name + " "
	}

	return
}
