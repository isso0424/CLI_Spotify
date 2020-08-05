package util

import (
	"isso0424/spotify_CLI/selfmadetypes/responseTypes"
	"math/rand"
	"time"
)

// Choose is function that get 1 item from slice.
func Choose(playlists []responseTypes.SearchResultItem) responseTypes.SearchResultItem {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(playlists))

	return playlists[index]
}

// CheckDuplicateName is function that check to exist playlist name.
func CheckDuplicateName(name string, playlistList []responseTypes.SearchResultItem) bool {
	for _, content := range playlistList {
		if content.Name == name {
			return false
		}
	}

	return true
}

// ExistTarget is function that judge include target in judgeTargets.
func ExistTarget(target string, judgeTargets []string) bool {
	for _, judgeTarget := range judgeTargets {
		if judgeTarget == target {
			return true
		}
	}

	return false
}
