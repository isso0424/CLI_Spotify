package command

import (
	"fmt"
	"net/http"
)

func createRequest(token string, method string, url string) (response *http.Response, err error) {
  request, err := http.NewRequest(method, url, nil)
  if err != nil {
    return
  }

  request.Header.Set("Authorization", "Bearer " + token)
  client := &http.Client{}
  response, err = client.Do(request)
  if err != nil {
    return
  }

  if response.StatusCode == 401 {
    fmt.Println("token is invalid\nyou have to execute `refresh` command")
  }

  return
}
