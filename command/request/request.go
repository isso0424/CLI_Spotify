package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"isso0424/spotify_CLI/auth"
	"isso0424/spotify_CLI/selfMadeTypes"
	"net/http"
)

const (
  baseURL = "https://api.spotify.com/v1"
  noContent = 204
  unAuthorized = 401
)

func CreateRequest(token *string, method selfMadeTypes.Method, requestUrl string, body io.Reader) (response *http.Response, err error) {
	request, err := http.NewRequest(method.String(), baseURL+requestUrl, body)
	if err != nil {
		return
	}

	request.Header.Set("Authorization", "Bearer "+*token)
	client := &http.Client{}
	response, err = client.Do(request)
	if err != nil {
		return
	}

	if response.StatusCode == unAuthorized {
		newTokenPtr, Err := auth.GetToken()
		if err != nil {
			err = Err
			return
		}

		*token = *newTokenPtr
	}

	return
}

func GetPlayListStatus(token *string, playlistID *string) (status selfMadeTypes.PlayListFromRequest, err error) {
	response, err := CreateRequest(token, selfMadeTypes.GET, fmt.Sprintf("/playlists/%s?fields=name%%2Cowner", *playlistID), nil)
	if err != nil {
		return
	}

	buffer := make([]byte, 1024)
	_, err = response.Body.Read(buffer)
	if err != nil {
		return
	}

	buffer = bytes.Trim(buffer, "\x00")

	err = json.Unmarshal(buffer, &status)
	if err != nil {
		return
	}

	return
}

func GetStatus(token *string) (status *selfMadeTypes.Content, err error) {
	response, err := CreateRequest(token, selfMadeTypes.GET, "/me/player", nil)
	if err != nil {
		return
	}
	if response.StatusCode == noContent {
		err = &selfMadeTypes.FailedGetError{Target: "playing status"}
		return
	}

	buffer := make([]byte, 8192)
	_, err = response.Body.Read(buffer)
	if err != nil {
		return
	}

	buffer = bytes.Trim(buffer, "\x00")

	err = json.Unmarshal(buffer, &status)

	return
}
