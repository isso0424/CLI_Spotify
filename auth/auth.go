package auth

import (
	"io/ioutil"
	"os"
)

const tokenFile = ".token"

func GetToken(force bool) (string, error) {
  if existDotToken() && !force {
    return readDotToken()
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

