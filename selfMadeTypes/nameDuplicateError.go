package selfMadeTypes

import "fmt"

type NameDuplicateError struct {
	Target string
}

func (ptr *NameDuplicateError) Error() string {
	return fmt.Sprintf("Already exist name %s", ptr.Target)
}
