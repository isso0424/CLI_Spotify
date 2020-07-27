package parse

import (
	"isso0424/spotify_CLI/selfmadetypes"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestPausingStatus is test function for CreatePlayinStatus
func TestPausingStatus(t *testing.T) {
	content := selfmadetypes.Content{
		IsPlaying: false,
	}

	playList := selfmadetypes.PlayListFromRequest{}

	assert.Equal(t, CreatePlayingStatus(content, playList), "Pausing")
}

// TestPlayingStatus is test function for CreatePlayingStatus
func TestPlayingStatus(t *testing.T) {
	content := selfmadetypes.Content{
		IsPlaying: true,
		Item: selfmadetypes.Item{
			Name: "name",
			Artists: []selfmadetypes.Artists{
				{
					Name: "artist",
				},
			},
		},
	}

	playList := selfmadetypes.PlayListFromRequest{
		Name: "playList",
		Owner: selfmadetypes.User{
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

// TestCreateContextURISuccess is test function for CreateContextURI
func TestCreateContextURISuccess(t *testing.T) {
	url := "https://open.spotify.com/playlist/37i9dQZF1DXd8cPo2t5Hqf?si=X4SkTg0BTHKclOIlM0D8lA"
	uri, _ := CreateContextURI(url)
	assert.Equal(t, *uri, "spotify:playlist:37i9dQZF1DXd8cPo2t5Hqf")
}

// TestCreateContextURIFailed is test functon for CreateContextURI
func TestCreateContextURIFailed(t *testing.T) {
	url := "https://open.spotify.com/"
	_, err := CreateContextURI(url)
	assert.EqualError(t, err, "too short length")
}

// TestGetPlaylistIDSuccess is test function.
func TestGetPlaylistIDSuccess(t *testing.T) {
	url := "https://open.spotify.com/playlist/37i9dQZF1DXd8cPo2t5Hqf?si=X4SkTg0BTHKclOIlM0D8lA"
	uri, _ := GetPlaylistID(url)
	assert.Equal(t, *uri, "37i9dQZF1DXd8cPo2t5Hqf")
}

// TestGetPlaylistIDFailed is test function.
func TestGetPlaylistIDFailed(t *testing.T) {
	url := "https://open.spotify.com/"
	_, err := GetPlaylistID(url)
	assert.EqualError(t, err, "too short length")
}

// TestLengthError is test function.
func TestLengthError(t *testing.T) {
	err := &lengthError{}
	assert.EqualError(t, err, "too short length")
}
