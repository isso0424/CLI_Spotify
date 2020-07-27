package selfmadetypes

import "fmt"

// FailedGetError is failed get error.
type FailedGetError struct {
	Target string
}

func (ptr *FailedGetError) Error() string {
	return fmt.Sprintf("Failed to get %s", ptr.Target)
}
