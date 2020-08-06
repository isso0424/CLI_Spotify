package parse

import (
	"fmt"
	"strings"
)

const urlMinimamLength = 5

// CreateContextURI is parse context uri function.
func CreateContextURI(url string) (string, error) {
	err := &lengthError{}
	spritted := strings.Split(url, "/")

	if len(spritted) < urlMinimamLength {
		return "", err
	}
	kind := spritted[3]
	tmp := spritted[4]

	id := strings.Split(tmp, "?")[0]

	contextURI := fmt.Sprintf("spotify:%s:%s", kind, id)
	return contextURI, nil
}

// GetIDFromURL is function that get playlist ID from url.
func GetIDFromURL(url string) (string, error) {
	err := &lengthError{}
	spritted := strings.Split(url, "/")

	if len(spritted) < urlMinimamLength {
		return "", err
	}
	tmp := spritted[4]

	id := strings.Split(tmp, "?")[0]

	return id, nil
}

// GetKindFromURL get URL kind.
func GetKindFromURL(url string) (string, error) {
	err := &lengthError{}
	splitted := strings.Split(url, "/")

	if len(splitted) < urlMinimamLength {
		return "", err
	}
	kind := splitted[3]

	return kind, nil
}
