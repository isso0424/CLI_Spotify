package file

import (
	"encoding/json"
	"errors"
	"isso0424/spotify_CLI/selfmadetypes"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestLoadPlayListSuccess is test function
func TestLoadPlayListSuccess(t *testing.T) {
	reset := setLoadPlayList(
		func(fileName string) (os.FileInfo, error) {
			fileInfo := files{}
			return fileInfo, nil
		},
		func(fileName string) ([]byte, error) {
			playlistList := []selfmadetypes.SearchResultItem{
				{
					Name: "SearchResultItem",
					URI:  "URI",
				},
			}
			return json.Marshal(playlistList)
		},
	)
	defer reset()

	successResult, _ := LoadPlayList()
	assert.Equal(
		t,
		[]selfmadetypes.SearchResultItem{
			{
				Name: "SearchResultItem",
				URI:  "URI",
			},
		},
		successResult,
	)
}

// TestLoadPlayListFail is test function
func TestLoadPlayListFail(t *testing.T) {
	reset := setLoadPlayList(
		func(fileName string) (os.FileInfo, error) {
			return nil, errors.New("file not exist")
		},
		func(fileName string) ([]byte, error) {
			playlistList := []selfmadetypes.SearchResultItem{
				{
					Name: "SearchResultItem",
					URI:  "URI",
				},
			}
			return json.Marshal(playlistList)
		},
	)

	_, err := LoadPlayList()
	assert.EqualError(t, err, "file not exist")
	reset()

	reset = setLoadPlayList(
		func(fileName string) (os.FileInfo, error) {
			fileInfo := files{}
			return fileInfo, nil
		},
		func(fileName string) ([]byte, error) {
			return nil, errors.New("cannot read file")
		},
	)

	_, err = LoadPlayList()
	assert.EqualError(t, err, "cannot read file")
	reset()
}

// TestSavePlayListSuccess is test function
func TestSavePlayListSuccess(t *testing.T) {
	reset := setSavePlayList(
		func(fileName string, fileDetail []byte, permission os.FileMode) error {
			return nil
		},
		func() ([]selfmadetypes.SearchResultItem, error) {
			playlistList := []selfmadetypes.SearchResultItem{
				{
					Name: "SearchResultItem",
					URI:  "URI",
				},
			}
			return playlistList, nil
		},
	)
	defer reset()

	err := SavePlayList(
		selfmadetypes.SearchResultItem{
			Name: "PlayList2",
			URI:  "URI",
		},
	)
	assert.Equal(t, nil, err)
}

// TestSavePlayListFail is test function
func TestSavePlayListFail(t *testing.T) {
	reset := setSavePlayList(
		func(fileName string, fileDetail []byte, permission os.FileMode) error {
			return errors.New("cannot write a file")
		},
		func() ([]selfmadetypes.SearchResultItem, error) {
			playlistList := []selfmadetypes.SearchResultItem{
				{
					Name: "SearchResultItem",
					URI:  "URI",
				},
			}
			return playlistList, nil
		},
	)

	err := SavePlayList(
		selfmadetypes.SearchResultItem{
			Name: "PlayList2",
			URI:  "URI",
		},
	)
	assert.EqualError(t, err, "cannot write a file")
	reset()

	reset = setSavePlayList(
		func(fileName string, fileDetail []byte, permission os.FileMode) error {
			return nil
		},
		func() ([]selfmadetypes.SearchResultItem, error) {
			return nil, errors.New("cannot load a file")
		},
	)

	err = SavePlayList(
		selfmadetypes.SearchResultItem{
			Name: "PlayList2",
			URI:  "URI",
		},
	)
	assert.EqualError(t, err, "cannot load a file")
	reset()
}

type files struct {
}

func (f files) Name() string {
	return "file"
}

func (f files) Size() int64 {
	return 1
}

func (f files) Mode() os.FileMode {
	const permission os.FileMode = 0600
	return permission
}

func (f files) ModTime() time.Time {
	return time.Time{}
}

func (f files) IsDir() bool {
	return false
}

func (f files) Sys() interface{} {
	return ""
}
