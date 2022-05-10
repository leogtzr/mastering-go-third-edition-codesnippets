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

func initS(N, S string, Y int) Entry {
	if Y < 2000 {
		return Entry{Name: N, Surname: S, Year: 2000}
	}

	return Entry{Name: N, Surname: S, Year: Y}
}

func main() {
	x := zeroS()
	fmt.Println(x)
}
