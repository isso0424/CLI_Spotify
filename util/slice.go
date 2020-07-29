package util

import (
	"isso0424/spotify_CLI/selfmadetypes"
	"math/rand"
	"time"
)

func Choose(playlists []selfmadetypes.PlayList) selfmadetypes.PlayList {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(playlists))

	return playlists[index]
}
