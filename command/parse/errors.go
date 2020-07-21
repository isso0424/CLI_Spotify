package parse

func (e *lengthError) Error() string {
	return "too short length"
}

type lengthError struct {
}
