// Package responseTypes is types for responseTypes
package responseTypes

import (
	"bytes"
	"net/http"
)

// Response is interafce for HTTP requestTypes responseTypes.
type Response interface {
	GetBody() []byte
	GetStatusCode() int
}

// HTTPResponse is struct for HTTP requestTypes responseTypes.
type HTTPResponse struct {
	Body       []byte
	StatusCode int
}

// GetBody get body.
func (response HTTPResponse) GetBody() []byte {
	return response.Body
}

// GetStatusCode get HTTP status code.
func (response HTTPResponse) GetStatusCode() int {
	return response.StatusCode
}

// New create new HTTPResponse
func (response HTTPResponse) New(resource *http.Response) (HTTPResponse, error) {
	statusCode := resource.StatusCode

	responseArray := make([]byte, 32768)
	_, err := resource.Body.Read(responseArray)
	if err != nil {
		return HTTPResponse{}, err
	}

	responseArray = bytes.Trim(responseArray, "\x00")

	return HTTPResponse{
		Body:       responseArray,
		StatusCode: statusCode,
	}, nil
}
