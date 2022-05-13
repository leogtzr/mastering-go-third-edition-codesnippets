package main

import (
	"fmt"
	"os"
	"regexp"
)

func matchNameSur(s string, re *regexp.Regexp) bool {
	t := []byte(s)

	return re.Match(t)
}

func main() {
	re := regexp.MustCompile(`^[A-Z][a-z]*$`)

	args := os.Args[1:]

	for _, arg := range args {
		fmt.Printf("%s -> %t\n", arg, matchNameSur(arg, re))
	}
}
