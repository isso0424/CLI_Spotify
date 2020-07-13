package selfMadeTypes

import "fmt"

type FailedGetError struct {
	Target string
}

func (ptr *FailedGetError) Error() string {
	return fmt.Sprintf("Failed to get %s", ptr.Target)
}
