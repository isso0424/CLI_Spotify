// Package requestTypes is type for requestTypes
package requestTypes

// Method is HTTP Request Methods wrapper
type Method int

const (
	// GET is GET method
	GET Method = iota
	// POST is POST method
	POST
	// PUT is PUT method
	PUT
	// DELETE is DELETE method
	DELETE
)

func (m Method) String() string {
	switch m {
	case GET:
		return "GET"
	case POST:
		return "POST"
	case PUT:
		return "PUT"
	case DELETE:
		return "DELETE"
	default:
		return "INVALID"
	}
}
