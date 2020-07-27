package file

import (
	"encoding/json"
	"io/ioutil"
	"isso0424/spotify_CLI/selfMadeTypes"
	"os"
)

var (
	fileExist func(string) (os.FileInfo, error)
	readFile  func(string) ([]byte, error)
)

func init() {
	fileExist = func(fileName string) (os.FileInfo, error) {
		return os.Stat(fileName)
	}

	readFile = func(fileName string) ([]byte, error) {
		return ioutil.ReadFile(fileName)
	}
}

func LoadPlayList() (playlistList []selfMadeTypes.PlayList, err error) {
	if _, err = fileExist("playlist.json"); err != nil {
		return
	}

	file, err := readFile("playlist.json")
	if err != nil {
		return
	}

	err = json.Unmarshal(file, &playlistList)

	return
}

func setLoadPlayList(
	fileExistFunc func(string) (os.FileInfo, error),
	readFileFunc func(string) ([]byte, error),
) func() {
	tmpFileExist := fileExist
	tmpReadFile := readFile
	fileExist = fileExistFunc
	readFile = readFileFunc

	return func() {
		fileExist = tmpFileExist
		readFile = tmpReadFile
	}
}
