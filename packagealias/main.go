package main

import (
	"fmt"
	s "strings"
	// "unicode"
)

var f = fmt.Printf

func main() {
	f(s.ToUpper("Leonardo"))
	f("\n")
	f("%t\n", s.EqualFold("Leo", "LEO"))
}
