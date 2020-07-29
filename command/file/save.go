package file

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"isso0424/spotify_CLI/selfmadetypes"
	"os"
)

var (
	writeFile func(string, []byte, os.FileMode) error
	loadFile  func() (playlistList []selfmadetypes.SearchResultItem, err error)
)

func init() {
	writeFile = func(fileName string, fileDetail []byte, permission os.FileMode) error {
		return ioutil.WriteFile(fileName, fileDetail, permission)
	}

	loadFile = func() (playlistList []selfmadetypes.SearchResultItem, err error) {
		return LoadPlayList()
	}
}

// SavePlayList is function save playlist for json file
func SavePlayList(target selfmadetypes.SearchResultItem) (err error) {
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

	fmt.Printf("\nplaylist saved!!!\nurl: %s\nname: %s\n", target.URI, target.Name)

	return
}

func setSavePlayList(
	writeFileFunc func(string, []byte, os.FileMode) error,
	loadFileFunc func() ([]selfmadetypes.SearchResultItem, error),
) func() {
	tmpWriteFile := writeFile
	tmpLoadFile := loadFile
	writeFile = writeFileFunc
	loadFile = loadFileFunc

	return func() {
		writeFile = tmpWriteFile
		loadFile = tmpLoadFile
	}
}
