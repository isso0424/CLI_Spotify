package file

import (
	"encoding/json"
	"errors"
	"isso0424/spotify_CLI/selfMadeTypes"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLoadPlayListSuccess(t *testing.T) {
  reset := setLoadPlayList(
    func(fileName string) (os.FileInfo, error) {
      fileInfo := files{}
      return fileInfo, nil
    },
    func(fileName string) ([]byte, error) {
      playlistList := []selfMadeTypes.PlayList{
        {
          Name: "PlayList",
          Uri: "URI",
        },
      }
	    return json.Marshal(playlistList)
    },
  )
  defer reset()

  successResult, _ := LoadPlayList()
  assert.Equal(
    t,
    []selfMadeTypes.PlayList{
      {
        Name: "PlayList",
        Uri: "URI",
      },
    },
    successResult,
  )
}

func TestLoadPlayListFail(t *testing.T) {
  reset := setLoadPlayList(
    func(fileName string) (os.FileInfo, error) {
      return nil, errors.New("file not exist")
    },
    func(fileName string) ([]byte, error) {
      playlistList := []selfMadeTypes.PlayList{
        {
          Name: "PlayList",
          Uri: "URI",
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

func TestSavePlayList(t *testing.T) {
  reset := setSavePlayList(
    func(fileName string, fileDetail []byte, permission os.FileMode) error {
      return nil
    },
    func() ([]selfMadeTypes.PlayList, error) {
      playlistList := []selfMadeTypes.PlayList{
        {
          Name: "PlayList",
          Uri: "URI",
        },
      }
      return playlistList, nil
    },
  )
  defer reset()

  err := SavePlayList(
    selfMadeTypes.PlayList{
      Name: "PlayList2",
      Uri: "URI",
    },
  )
  assert.Equal(t, nil, err)
}

type files struct {
}

func(f files) Name() string {
  return "file"
}

func(f files) Size() int64 {
  return 1
}

func(f files) Mode() os.FileMode{
  return 0600
}

func(f files) ModTime() time.Time {
  return time.Time{}
}

func(f files) IsDir() bool {
  return false
}

func(f files) Sys() interface{} {
  return ""
}
