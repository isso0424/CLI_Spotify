package parse

import (
	"fmt"
	"strings"
)

func CreateContextUri(url string) (*string, error) {
	err := &lengthError{}
	spritted := strings.Split(url, "/")

	if len(spritted) < 5 {
		return nil, err
	}
	kind := spritted[3]
	tmp := spritted[4]

	id := strings.Split(tmp, "?")[0]

	context_uri := fmt.Sprintf("spotify:%s:%s", kind, id)
	return &context_uri, nil
}

