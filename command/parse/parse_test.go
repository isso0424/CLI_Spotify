package parse

import (
	"isso0424/spotify_CLI/selfmadetypes"
	"isso0424/spotify_CLI/selfmadetypes/response"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestPausingStatus is test function for CreatePlayinStatus
func TestPausingStatus(t *testing.T) {
	content := response.Content{
		IsPlaying: false,
	}

	playList := response.PlayList{}

	assert.Equal(
		t,
		CreatePlayingStatus(content, playList.Name, playList.Owner.DisplayName, "playlist"),
		selfmadetypes.OutputMessage{Message: [][]string{{"Pausing"}}},
	)
}

// TestPlayingStatus is test function for CreatePlayingStatus
func TestPlayingStatus(t *testing.T) {
	content := response.Content{
		IsPlaying: true,
		Item: response.Item{
			Name: "name",
			Artists: []response.Artists{
				{
					Name: "artist",
				},
			},
		},
	}

	playList := response.PlayList{
		Name: "playList",
		Owner: response.User{
			DisplayName: "user",
		},
	}

	assert.Equal(
		t,
		CreatePlayingStatus(content, playList.Name, playList.Owner.DisplayName, "playlist"),
		selfmadetypes.OutputMessage{
			Message: [][]string{
				{
					"Playing status",
				},
				{
					"Title: name",
					"Artist: artist",
				},
				{
					"Playing playlist",
				},
				{
					"Playlist name: playList",
					"Owner: user",
				},
			},
		},
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
	uri, _ := GetIDFromURL(url)
	assert.Equal(t, *uri, "37i9dQZF1DXd8cPo2t5Hqf")
}

// TestGetPlaylistIDFailed is test function.
func TestGetPlaylistIDFailed(t *testing.T) {
	url := "https://open.spotify.com/"
	_, err := GetIDFromURL(url)
	assert.EqualError(t, err, "too short length")
}

// TestLengthError is test function.
func TestLengthError(t *testing.T) {
	err := &lengthError{}
	assert.EqualError(t, err, "too short length")
}
