package command

import (
	"isso0424/spotify_CLI/selfmadetypes"
	"isso0424/spotify_CLI/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSwitchRepeat is test function for switchRepeat()
func TestSwitchRepeat(t *testing.T) {
	var result string
	result = util.SwitchRepeatState("off")
	assert.Equal(t, result, "context")

	result = util.SwitchRepeatState("track")
	assert.Equal(t, result, "off")

	result = util.SwitchRepeatState("context")
	assert.Equal(t, result, "track")
}

// TestCheckDuplicateName is test function for CheckDuplicateName()
func TestCheckDuplicateName(t *testing.T) {
	var playlistList []selfmadetypes.SearchResultItem
	var result bool
	playlistList = []selfmadetypes.SearchResultItem{
		{
			Name: "playlist",
			URI:  "hogefuga",
		},
		{
			Name: "playlist2",
			URI:  "unchi",
		},
	}

	result = util.CheckDuplicateName("playlist", playlistList)
	assert.Equal(t, result, false)

	result = util.CheckDuplicateName("not found", playlistList)
	assert.Equal(t, result, true)
}
