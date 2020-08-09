package util

import (
	"isso0424/spotify_CLI/selfmadetypes/responsetypes"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSwitchRepeat is test function for switchRepeat()
func TestSwitchRepeat(t *testing.T) {
	var result string
	result = SwitchRepeatState("off")
	assert.Equal(t, result, "context")

	result = SwitchRepeatState("track")
	assert.Equal(t, result, "off")

	result = SwitchRepeatState("context")
	assert.Equal(t, result, "track")
}

// TestCheckDuplicateName is test function for CheckDuplicateName()
func TestCheckDuplicateName(t *testing.T) {
	var playlistList []responsetypes.SearchResultItem
	var result bool
	playlistList = []responsetypes.SearchResultItem{
		{
			Name: "playlist",
			URI:  "hogefuga",
		},
		{
			Name: "playlist2",
			URI:  "unchi",
		},
	}

	result = CheckDuplicateName("playlist", playlistList)
	assert.Equal(t, result, false)

	result = CheckDuplicateName("not found", playlistList)
	assert.Equal(t, result, true)
}
