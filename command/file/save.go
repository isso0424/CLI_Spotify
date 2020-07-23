package file

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"isso0424/spotify_CLI/selfMadeTypes"
	"os"
)

var (
  writeFile func(string, []byte, os.FileMode) error
  loadFile func() (playlistList []selfMadeTypes.PlayList, err error)
)

func init() {
  writeFile = func(fileName string, fileDetail []byte, permission os.FileMode) error {
    return ioutil.WriteFile(fileName, fileDetail, permission)
  }

  loadFile = func() (playlistList []selfMadeTypes.PlayList, err error) {
    return LoadPlayList()
  }
}

func SavePlayList(target selfMadeTypes.PlayList) (err error) {
	playlistList, err := loadFile()
  if err != nil {
    return
  }
	playlistList = append(playlistList, target)

	jsonFile, err := json.Marshal(playlistList)
	if err != nil {
		return
	}

	err = writeFile("playlist.json", jsonFile, 0666)

	if err != nil {
		return
	}

	fmt.Printf("\nplaylist saved!!!\nurl: %s\nname: %s\n", target.Uri, target.Name)

	return
}

func setSavePlayList(writeFileFunc func(string, []byte, os.FileMode) error, loadFileFunc func() ([]selfMadeTypes.PlayList, error)) func() {
  tmpWriteFile := writeFile
  tmpLoadFile := loadFile
  writeFile = writeFileFunc
  loadFile = loadFileFunc

  return func() {
    writeFile = tmpWriteFile
    loadFile = tmpLoadFile
  }
}
