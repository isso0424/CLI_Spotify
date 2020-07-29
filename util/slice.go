package util

import (
	"isso0424/spotify_CLI/selfmadetypes"
	"math/rand"
	"time"
)

func Choose(playlists []selfmadetypes.SearchResultItem) selfmadetypes.SearchResultItem {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(playlists))

	return playlists[index]
}

func CheckDuplicateName(name string, playlistList []selfmadetypes.SearchResultItem) bool {
	for _, content := range playlistList {
		if content.Name == name {
			return false
		}
	}

	return true
}

func ExistTarget(target string, judgeTargets []string) bool {
	for _, judgeTarget := range judgeTargets {
		if judgeTarget == target {
			return true
		}
	}

	return false
}
