// Package util is utility for command
package util

import (
	"bufio"
	"fmt"
	"os"
)

// Input is get stdinput string.
func Input(message string, hint string) (inputtedValue string) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("%s\n%s|>>>", message, hint)
	if scanner.Scan() {
		inputtedValue = scanner.Text()
	}
	fmt.Println()

	return
}
