package selfMadeTypes

type Method int

const (
	GET Method = iota
	POST
	PUT
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
