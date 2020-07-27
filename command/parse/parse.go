package parse

import (
	"isso0424/spotify_CLI/selfMadeTypes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPausingStatus(t *testing.T) {
	content := selfMadeTypes.Content{
		IsPlaying: false,
	}

	playList := selfMadeTypes.PlayListFromRequest{}

	assert.Equal(t, CreatePlayingStatus(content, playList), "Pausing")
}

func TestPlayingStatus(t *testing.T) {
	content := selfMadeTypes.Content{
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

	assert.Equal(
		t,
		CreatePlayingStatus(content, playList),
		"Playing status\n"+
			"--------------\n"+
			"Title: name\n"+
			"Artist: artist\n\n"+
			"PlayList Information\n"+
			"-------------------\n"+
			"PlayList: playList\n"+
			"Owner: user\n",
	)
}

func TestCreateContextUriSuccess(t *testing.T) {
	url := "https://open.spotify.com/playlist/37i9dQZF1DXd8cPo2t5Hqf?si=X4SkTg0BTHKclOIlM0D8lA"
	uri, _ := CreateContextUri(url)
	assert.Equal(t, *uri, "spotify:playlist:37i9dQZF1DXd8cPo2t5Hqf")
}

func TestCreateContextUriFailed(t *testing.T) {
	url := "https://open.spotify.com/"
	_, err := CreateContextUri(url)
	assert.EqualError(t, err, "too short length")
}

func TestGetPlaylistIDSuccess(t *testing.T) {
	url := "https://open.spotify.com/playlist/37i9dQZF1DXd8cPo2t5Hqf?si=X4SkTg0BTHKclOIlM0D8lA"
	uri, _ := GetPlaylistID(url)
	assert.Equal(t, *uri, "37i9dQZF1DXd8cPo2t5Hqf")
}

func TestGetPlaylistIDFailed(t *testing.T) {
	url := "https://open.spotify.com/"
	_, err := GetPlaylistID(url)
	assert.EqualError(t, err, "too short length")
}

func TestLengthError(t *testing.T) {
	err := &lengthError{}
	assert.EqualError(t, err, "too short length")
}
