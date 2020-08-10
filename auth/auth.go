// Package auth is authrize with spotify package.
package auth

import (
	"fmt"
	"io/ioutil"
	"os"
)

var(
	tokenFile string
	configFile string
)

func init() {
	configDir := os.Getenv("XDG_CONFIG_HOME")
	if configDir == "" {
		configDir = os.Getenv("HOME")
	}
	tokenFile = fmt.Sprintf("%s/.config/spotify_CLI/.token", configDir)
	configFile = fmt.Sprintf("%s/.config/spotify_CLI/config", configDir)
}

// GetToken is function that get token from .token.
func GetToken() (*string, error) {
	if existDotToken() {
		token, err := readDotToken()
		if err != nil {
			return nil, err
		}
		return refresh(token)
	}

	return oauth()
}

func existDotToken() bool {
	_, err := os.Stat(tokenFile)
	return err == nil
}

func readDotToken() (string, error) {
	bytes, err := ioutil.ReadFile(tokenFile)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
