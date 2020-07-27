package command

import (
	"isso0424/spotify_CLI/selfmadetypes"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSwitchRepeat is test function for switchRepeat()
func TestSwitchRepeat(t *testing.T) {
	var result string
	result = switchRepeatState("off")
	assert.Equal(t, result, "context")

	result = switchRepeatState("track")
	assert.Equal(t, result, "off")

	result = switchRepeatState("context")
	assert.Equal(t, result, "track")
}

// TestCheckDuplicateName is test function for checkDuplicateName()
func TestCheckDuplicateName(t *testing.T) {
	var playlistList []selfmadetypes.PlayList
	var result bool
	playlistList = []selfmadetypes.PlayList{
		{
			Name: "playlist",
			URI:  "hogefuga",
		},
		{
			Name: "playlist2",
			URI:  "unchi",
		},
	}

	result = checkDuplicateName("playlist", playlistList)
	assert.Equal(t, result, false)

	result = checkDuplicateName("not found", playlistList)
	assert.Equal(t, result, true)
}
