package parse

import (
	"isso0424/spotify_CLI/selfMadeTypes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPausingStatus(t *testing.T) {
  content := selfMadeTypes.Content {
    IsPlaying: false,
  }

  playList := selfMadeTypes.PlayListFromRequest{}

  assert.Equal(t, CreatePlayingStatus(content, playList), "Pausing")
}

func TestPlayingStatus(t *testing.T) {
  content := selfMadeTypes.Content {
    IsPlaying: true,
    Item: selfMadeTypes.Item{
      Name: "name",
      Artists: []selfMadeTypes.Artists{
        {
          Name: "artist",
        },
      },
    },
  }

  playList := selfMadeTypes.PlayListFromRequest{
    Name: "playList",
    Owner: selfMadeTypes.User{
      DisplayName: "user",
    },
  }

  assert.Equal(t, CreatePlayingStatus(content, playList), "Playing status\n--------------\nTitle: name\nArtist: artist\n\nPlayList Infomation\n-------------------\nPlayList: playList\nOwner: user\n")
}
