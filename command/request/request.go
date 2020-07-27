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
	baseURL      = "https://api.spotify.com/v1"
	noContent    = 204
	unAuthorized = 401
)

func CreateRequest(token *string, method selfMadeTypes.Method, requestUrl string, body io.Reader) (responseArray []byte, err error, statusCode int) {
	request, err := http.NewRequest(method.String(), baseURL+requestUrl, body)
	if err != nil {
		return
	}

	request.Header.Set("Authorization", "Bearer "+*token)
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()

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

	return responseArray, err, statusCode
}

func GetPlayListStatus(token *string, playlistID *string) (status selfMadeTypes.PlayListFromRequest, err error) {
	response, err, _ := CreateRequest(
		token,
		selfMadeTypes.GET,
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

func GetStatus(token *string) (status *selfMadeTypes.Content, err error) {
	response, err, statusCode := CreateRequest(token, selfMadeTypes.GET, "/me/player", nil)
	if err != nil {
		return
	}
	if statusCode == noContent {
		err = &selfMadeTypes.FailedGetError{Target: "playing status"}
		return
	}

	err = json.Unmarshal(response, &status)

	return
}
