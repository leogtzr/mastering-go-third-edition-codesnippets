package main

import (
	"fmt"
	"unicode"
)

func main() {
	// name := "Leonardo Gutierrez Ramirez"
	name := "\x99\x00ab\x50\x00\x23\x50\x29\x9c"
	for i := 0; i < len(name); i++ {
		if unicode.IsPrint(rune(name[i])) {
			fmt.Printf("%c\n", name[i])
		} else {
			fmt.Println("Not printable!")
		}
	}
}
