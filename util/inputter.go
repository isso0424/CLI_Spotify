package util

import (
	"bufio"
	"fmt"
	"os"
)

// Input is get stdinput string.
func Input(message string, hint string, getValiable *string) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("%s\n%s|>>>", message, hint)
	if scanner.Scan() {
		*getValiable = scanner.Text()
	}
	fmt.Println()
}
