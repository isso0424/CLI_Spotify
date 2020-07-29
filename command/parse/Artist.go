package parse

import "isso0424/spotify_CLI/selfmadetypes"

// GetArtistNames is extract artist names from artists object slice.
func GetArtistNames(artists []selfmadetypes.Artists) (artistNames string) {
	artistNames = ""

	for _, artist := range artists {
		artistNames += artist.Name + " "
	}

	return
}
