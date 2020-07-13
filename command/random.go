package command

import (
	"isso0424/spotify_CLI/selfMadeTypes"
	"isso0424/spotify_CLI/util"
	"math/rand"
	"time"
)

func random(token string) (newToken string, err error) {
	newToken = token
	playlists, err := util.LoadPlayList()
	if err != nil {
		return
	}

	playlist := choice(playlists)

	newToken, err = play(token, playlist.Uri)

	return
}

func choice(playlists []selfMadeTypes.PlayList) selfMadeTypes.PlayList {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(playlists))

	return playlists[index]
}
