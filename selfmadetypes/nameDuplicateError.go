package selfmadetypes

import "fmt"

// NameDuplicateError is name duplicate error.
type NameDuplicateError struct {
	Target string
}

func (ptr *NameDuplicateError) Error() string {
	return fmt.Sprintf("Already exist name %s", ptr.Target)
}
