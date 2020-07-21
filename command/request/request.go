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

func CreateRequest(token string, method string, url string, body io.Reader) (response *http.Response, newToken string, err error) {
	newToken = token
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return
	}

	request.Header.Set("Authorization", "Bearer "+token)
	client := &http.Client{}
	response, err = client.Do(request)
	if err != nil {
		return
	}

	if response.StatusCode == 401 {
		newTokenPtr, Err := auth.GetToken()
		if err != nil {
			err = Err
			return
		}

		newToken = *newTokenPtr
	}

	return
}

func GetPlayListStatus(token string, playlistID *string) (status selfMadeTypes.PlayListFromRequest, err error) {
		response, _, err := CreateRequest(token, "GET", fmt.Sprintf("https://api.spotify.com/v1/playlists/%s?fields=name%%2Cowner", *playlistID), nil)
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
