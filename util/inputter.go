package util

import "fmt"

func Input(hint string, getValiable *string) {
  fmt.Printf("\n%s|>>>", hint)
  fmt.Scanln(getValiable)
  fmt.Println()
}
