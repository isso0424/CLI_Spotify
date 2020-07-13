package command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"isso0424/spotify_CLI/selfMadeTypes"
	"isso0424/spotify_CLI/util"
	"strings"
)

func getPlayStatus(token string) (bool, string, error) {
	status, newToken, err := getStatus(token)
	if err != nil {
		return false, newToken, err
	}

	if status == nil {
		return false, newToken, nil
	}

  err = createInfo(*status, token)

  if err != nil {
    return false, newToken, err
  }

	return status.IsPlaying && len(status.Item.Artists) != 0, newToken, nil
}

func getStatus(token string) (status *selfMadeTypes.Content, newToken string, err error) {
	response, newToken, err := util.CreateRequest(token, "GET", "https://api.spotify.com/v1/me/player", nil)
	if err != nil {
		return
	}
	if response.StatusCode == 204 {
    err = &selfMadeTypes.FailedGetError{Target: "playing status"}
		return
	}

	buffer := make([]byte, 8192)
	_, err = response.Body.Read(buffer)

	buffer = bytes.Trim(buffer, "\x00")

	err = json.Unmarshal(buffer, &status)

	return
}

func createInfo(content selfMadeTypes.Content, token string) error {
	if content.IsPlaying && len(content.Item.Artists) != 0 {
    playlistUrl := content.Context.ExternalUrls.Spotify
    playlistID, err := getPlaylistID(playlistUrl)

    if err != nil {
      return err
    }

    response, _, err := util.CreateRequest(token, "GET", fmt.Sprintf("https://api.spotify.com/v1/playlists/%s?fields=name%%2Cowner", *playlistID), nil)
    if err != nil {
      return err
    }
    var playlist selfMadeTypes.PlayListFromRequest
	  buffer := make([]byte, 1024)
	  _, err = response.Body.Read(buffer)
    if err != nil {
      return err
    }

	  buffer = bytes.Trim(buffer, "\x00")

	  err = json.Unmarshal(buffer, &playlist)
    if err != nil {
      return err
    }

    fmt.Printf("Playing status\n--------------\nTitle: %s\nArtist: %s\n\nPlayList: %s\nOwner: %s\n", content.Item.Name, content.Item.Artists[0].Name, playlist.Name, playlist.Owner.DisplayName)
	} else {
		fmt.Println("Pausing")
	}

  return nil
}

func getPlaylistID(url string) (*string, error) {
	err := &lengthError{}
	spritted := strings.Split(url, "/")

	if len(spritted) < 5 {
    return nil, err
	}
	tmp := spritted[4]

	id := strings.Split(tmp, "?")[0]

  return &id, nil
}
