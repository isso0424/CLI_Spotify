package auth

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func refresh(token string) (newToken *string, err error) {
	if err != nil {
		return nil, err
	}

	form := url.Values{}
	form.Add("grant_type", "refresh_token")
	form.Add("refresh_token", token)

	request, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(form.Encode()))

	if err != nil {
		return
	}

	encoded := createEncodedID()

	request.Header.Set("Authorization", fmt.Sprintf("Basic %s", encoded))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return
	}
	defer func() {
		err = response.Body.Close()
		if err != nil {
			fmt.Println("Error: " + err.Error())
		}
	}()

	buffer := make([]byte, 1024)
	_, err = response.Body.Read(buffer)
	if err != nil {
		return
	}

	buffer = bytes.Trim(buffer, "\x00")

	var responseBody refreshTokenResponse
	err = json.Unmarshal(buffer, &responseBody)
	if err != nil {
		return
	}

	newToken = &responseBody.AccessToken

	return newToken, err
}

func createEncodedID() string {
	clientID, secretID := getClientID()
	ids := fmt.Sprintf("%s:%s", clientID, secretID)

	return base64.StdEncoding.EncodeToString([]byte(ids))
}

type refreshTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
	ExpiresIn   int32  `json:"expires_in"`
}
