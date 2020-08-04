// Package request is submit request package.
package request

import (
	"encoding/json"
	"fmt"
	"io"
	"isso0424/spotify_CLI/auth"
	"isso0424/spotify_CLI/selfmadetypes/request"
	"isso0424/spotify_CLI/selfmadetypes/response"
	"isso0424/spotify_CLI/selfmadetypes/selfmadeerrors"
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
	httpResponse response.Response,
	err error,
) {
	req, err := http.NewRequest(method.String(), baseURL+requestURL, body)
	if err != nil {
		return
	}

	req.Header.Set("Authorization", "Bearer "+*token)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return
	}
	defer func() {
		err := res.Body.Close()
		if err != nil {
			fmt.Printf("Error: " + err.Error())
		}
	}()

	httpResponse, err = response.HTTPResponse{}.New(res)

	if res.StatusCode == unAuthorized {
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
func GetPlayListStatus(token *string, playlistID *string) (status response.PlayList, err error) {
	res, err := CreateRequest(
		token,
		request.GET,
		fmt.Sprintf(
			"/playlists/%s?fields=name%%2Cowner",
			*playlistID,
		),
		nil,
	)
	if err != nil {
		return
	}

	err = json.Unmarshal(res.GetBody(), &status)
	if err != nil {
		return
	}

	return
}

// GetArtistStatus get artist status that is playing.
func GetArtistStatus(token *string, artistID *string) (status response.Artists, err error) {
	res, err := CreateRequest(
		token,
		request.GET,
		fmt.Sprintf(
			"/artists/%s",
			*artistID,
		),
		nil,
	)
	if err != nil {
		return
	}

	err = json.Unmarshal(res.GetBody(), &status)
	if err != nil {
		return
	}

	return
}

// GetAlbumStatus get album status that is playing.
func GetAlbumStatus(token *string, albumID *string) (status response.Album, err error) {
	res, err := CreateRequest(
		token,
		request.GET,
		fmt.Sprintf(
			"/albums/%s",
			*albumID,
		),
		nil,
	)
	if err != nil {
		return
	}

	err = json.Unmarshal(res.GetBody(), &status)
	if err != nil {
		return
	}

	return
}

// GetStatus is function that get playing status.
func GetStatus(token *string) (status *response.Content, err error) {
	res, err := CreateRequest(token, request.GET, "/me/player", nil)
	if err != nil {
		return
	}
	if res.GetStatusCode() == noContent {
		err = &command.FailedGetError{Target: "playing status"}
		return
	}

	err = json.Unmarshal(res.GetBody(), &status)

	return
}
