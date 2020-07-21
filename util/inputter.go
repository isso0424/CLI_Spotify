package util

import "fmt"

func Input(message string, hint string, getValiable *string) {
	fmt.Printf("%s\n%s|>>>", message, hint)
	fmt.Scanln(getValiable)
	fmt.Println()
}
