package main

import (
	"fmt"
	"os"
	"regexp"
)

func matchInteger(x string, re *regexp.Regexp) bool {
	input := []byte(x)

	return re.Match(input)
}

func main() {
	re := regexp.MustCompile(`^[-+]?\d+$`)

	args := os.Args[1:]
	for _, arg := range args {
		fmt.Printf("%s -> %t\n", arg, matchInteger(arg, re))
	}
}
