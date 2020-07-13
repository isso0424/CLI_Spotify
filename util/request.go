package util

import (
	"io"
	"isso0424/spotify_CLI/auth"
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
