package auth

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func refresh(token string) (newToken *string, err error) {
	err = godotenv.Load()
	if err != nil {
		return nil, err
	}

	form := url.Values{}
	form.Add("grant_type", "refresh_token")
	form.Add("refresh_token", token)

	request, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(form.Encode()))

	encoded := createEncodedID()

	request.Header.Set("Authorization", fmt.Sprintf("Basic %s", encoded))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return
	}

	buffer := make([]byte, 1024)
	_, err = response.Body.Read(buffer)
	if err != nil {
		return
	}

	buffer = bytes.Trim(buffer, "\x00")

	var responseBody refreshTokenResponse
	if err := json.Unmarshal(buffer, &responseBody); err != nil {
		fmt.Println("Error: ", err)
		return nil, err
	}

	newToken = &responseBody.AccessToken

	return
}

func createEncodedID() string {
	ids := fmt.Sprintf("%s:%s", os.Getenv("SPOTIFY_ID"), os.Getenv("SPOTIFY_SECRET"))

	return base64.StdEncoding.EncodeToString([]byte(ids))
}

type refreshTokenStruct struct {
	GrantType    string `json:"grant_type"`
	RefreshToken string `json:"refresh_token"`
}

type refreshTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
	ExpiresIn   int32  `json:"expires_in"`
}
