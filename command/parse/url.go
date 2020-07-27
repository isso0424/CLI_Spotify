package parse

import (
	"fmt"
	"strings"
)

const urlMinimamLength = 5

// CreateContextURI is parse context uri function.
func CreateContextURI(url string) (*string, error) {
	err := &lengthError{}
	spritted := strings.Split(url, "/")

	if len(spritted) < urlMinimamLength {
		return nil, err
	}
	kind := spritted[3]
	tmp := spritted[4]

	id := strings.Split(tmp, "?")[0]

	contextURI := fmt.Sprintf("spotify:%s:%s", kind, id)
	return &contextURI, nil
}

// GetPlaylistID is function that get playlist ID from url.
func GetPlaylistID(url string) (*string, error) {
	err := &lengthError{}
	spritted := strings.Split(url, "/")

	if len(spritted) < urlMinimamLength {
		return nil, err
	}
	tmp := spritted[4]

	id := strings.Split(tmp, "?")[0]

	return &id, nil
}
