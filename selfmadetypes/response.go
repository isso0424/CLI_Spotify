package selfmadetypes

import (
	"bytes"
	"net/http"
)

// Response is interafce for HTTP request response.
type Response interface {
	GetBody() []byte
	GetStatusCode() int
}

// HttpResponse is struct for HTTP request response.
type HttpResponse struct {
	Body       []byte
	StatusCode int
}

func (response HttpResponse) GetBody() []byte {
	return response.Body
}

func (response HttpResponse) GetStatusCode() int {
	return response.StatusCode
}

func (tmp HttpResponse) New(response *http.Response) (HttpResponse, error) {
	statusCode := response.StatusCode

	responseArray := make([]byte, 32768)
	_, err := response.Body.Read(responseArray)
	if err != nil {
		return HttpResponse{}, err
	}

	responseArray = bytes.Trim(responseArray, "\x00")

	return HttpResponse{
		Body:       responseArray,
		StatusCode: statusCode,
	}, nil
}
