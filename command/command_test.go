package command

import (
	"isso0424/spotify_CLI/selfMadeTypes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwitchRepeat(t *testing.T) {
  var result string
  result = switchRepeatState("off")
  assert.Equal(t, result, "context")

  result = switchRepeatState("track")
  assert.Equal(t, result, "off")

  result = switchRepeatState("context")
  assert.Equal(t, result, "track")
}

func TestCheckDuplicateName(t *testing.T) {
  var playlistList []selfMadeTypes.PlayList
  var result bool
  playlistList = []selfMadeTypes.PlayList{
    {
      Name: "playlist",
      Uri: "hogefuga",
    },
    {
      Name: "playlist2",
      Uri: "unchi",
    },
  }

  result = checkDuplicateName("playlist", playlistList)
  assert.Equal(t, result, false)

  result = checkDuplicateName("not found", playlistList)
  assert.Equal(t, result, true)
}
