package parse

// Error is error.
func (e *lengthError) Error() string {
	return "too short length"
}

type lengthError struct {
}
