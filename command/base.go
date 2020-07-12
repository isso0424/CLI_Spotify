package command

import (
	"isso0424/spotify_CLI/auth"
	"net/http"
)

func createRequest(token string, method string, url string) (response *http.Response, newToken string, err error) {
  newToken = token
	request, err := http.NewRequest(method, url, nil)
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
