package main

import (
	"fmt"
)

type Entry struct {
	Name string
	Surname string
	Year int
}

// Initialized by Go
func zeroS() Entry {
	return Entry{}
}


func main() {
	x := zeroS()
	fmt.Println(x)
}
