// Package commanderrors is error for command
package commanderrors

import "fmt"

// FailedGetError is failed get error.
type FailedGetError struct {
	Target string
}

func (ptr *FailedGetError) Error() string {
	return fmt.Sprintf("Failed to get %s", ptr.Target)
}

// NameDuplicateError is name duplicate error.
type NameDuplicateError struct {
	Target string
}

func (ptr *NameDuplicateError) Error() string {
	return fmt.Sprintf("Already exist name %s", ptr.Target)
}

// NotFound is not found error.
type NotFound struct {
	Target string
}

func (ptr *NotFound) Error() string {
	return fmt.Sprintf("%s is not found", ptr.Target)
}
