package command

import (
	"fmt"
	"isso0424/spotify_CLI/selfMadeTypes"
	"isso0424/spotify_CLI/util"
	"math/rand"
	"time"
)

func Random(token string) (newToken string) {
  newToken = token
	playlists, err := util.LoadPlayList()
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	playlist := choice(playlists)

	newToken = play(token, playlist.Uri)

  return
}

func choice(playlists []selfMadeTypes.PlayList) selfMadeTypes.PlayList {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(playlists))

	return playlists[index]
}
