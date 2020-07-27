package selfmadetypes

import "fmt"

// NotFound is not found error.
type NotFound struct {
	Target string
}

func (ptr *NotFound) Error() string {
	return fmt.Sprintf("%s is not found", ptr.Target)
}
