package parse

import (
	"isso0424/spotify_CLI/selfmadetypes/response"
)

// GetArtistNames is extract artist names from artists object slice.
func GetArtistNames(artists []response.Artists) (artistNames string) {
	artistNames = ""

	for _, artist := range artists {
		artistNames += artist.Name + " "
	}

	return
}
