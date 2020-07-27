package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"isso0424/spotify_CLI/auth"
	"isso0424/spotify_CLI/selfmadetypes"
	"net/http"
)

const (
	baseURL      = "https://api.spotify.com/v1"
	noContent    = 204
	unAuthorized = 401
)

// CreateRequest is new request and submit request function.
// Get response value.
func CreateRequest(token *string, method selfmadetypes.Method, requestURL string, body io.Reader) (responseArray []byte, statusCode int, err error) {
	request, err := http.NewRequest(method.String(), baseURL+requestURL, body)
	if err != nil {
		return
	}

	request.Header.Set("Authorization", "Bearer "+*token)
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return
	}
	defer func() {
		err := response.Body.Close()
		if err != nil {
			fmt.Printf("Error: " + err.Error())
		}
	}()

	statusCode = response.StatusCode

	responseArray = make([]byte, 32768)
	_, err = response.Body.Read(responseArray)
	if err != nil {
		return
	}

	responseArray = bytes.Trim(responseArray, "\x00")

	if response.StatusCode == unAuthorized {
		var newTokenPtr *string
		newTokenPtr, err = auth.GetToken()
		if err != nil {
			return
		}

		*token = *newTokenPtr
	}

	return responseArray, statusCode, err
}

// GetPlayListStatus is get user playlist status.
func GetPlayListStatus(token *string, playlistID *string) (status selfmadetypes.PlayListFromRequest, err error) {
	response, _, err := CreateRequest(
		token,
		selfmadetypes.GET,
		fmt.Sprintf(
			"/playlists/%s?fields=name%%2Cowner",
			*playlistID,
		),
		nil,
	)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &status)
	if err != nil {
		return
	}

	return
}

// GetStatus is function that get playing status.
func GetStatus(token *string) (status *selfmadetypes.Content, err error) {
	response, statusCode, err := CreateRequest(token, selfmadetypes.GET, "/me/player", nil)
	if err != nil {
		return
	}
	if statusCode == noContent {
		err = &selfmadetypes.FailedGetError{Target: "playing status"}
		return
	}

	err = json.Unmarshal(response, &status)

	return
}
