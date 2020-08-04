// Package request is submit request package.
package request

import (
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
func CreateRequest(
	token *string,
	method fmt.Stringer,
	requestURL string,
	body io.Reader,
) (
	httpResponse selfmadetypes.Response,
	err error,
) {
	request, err := http.NewRequest(method.String(), baseURL+requestURL, body)
	if err != nil {
		return
	}

	request.Header.Set("Authorization", "Bearer "+*token)
	request.Header.Set("Content-Type", "application/json")
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

	httpResponse, err = selfmadetypes.HTTPResponse{}.New(response)

	if response.StatusCode == unAuthorized {
		var newTokenPtr *string
		newTokenPtr, err = auth.GetToken()
		if err != nil {
			return
		}

		*token = *newTokenPtr
	}

	return httpResponse, err
}

// GetPlayListStatus is get user playlist status.
func GetPlayListStatus(token *string, playlistID *string) (status selfmadetypes.PlayList, err error) {
	response, err := CreateRequest(
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

	err = json.Unmarshal(response.GetBody(), &status)
	if err != nil {
		return
	}

	return
}

// GetArtistStatus get artist status that is playing.
func GetArtistStatus(token *string, artistID *string) (status selfmadetypes.Artists, err error) {
	response, err := CreateRequest(
		token,
		selfmadetypes.GET,
		fmt.Sprintf(
			"/artists/%s",
			*artistID,
		),
		nil,
	)
	if err != nil {
		return
	}

	err = json.Unmarshal(response.GetBody(), &status)
	if err != nil {
		return
	}

	return
}

// GetAlbumStatus get album status that is playing.
func GetAlbumStatus(token *string, albumID *string) (status selfmadetypes.Album, err error) {
	response, err := CreateRequest(
		token,
		selfmadetypes.GET,
		fmt.Sprintf(
			"/albums/%s",
			*albumID,
		),
		nil,
	)
	if err != nil {
		return
	}

	err = json.Unmarshal(response.GetBody(), &status)
	if err != nil {
		return
	}

	return
}

// GetStatus is function that get playing status.
func GetStatus(token *string) (status *selfmadetypes.Content, err error) {
	response, err := CreateRequest(token, selfmadetypes.GET, "/me/player", nil)
	if err != nil {
		return
	}
	if response.GetStatusCode() == noContent {
		err = &selfmadetypes.FailedGetError{Target: "playing status"}
		return
	}

	err = json.Unmarshal(response.GetBody(), &status)

	return
}
