package main

import (
	"fmt"
)

type Entry struct {
	Name string
	Surname string
	Year int
}

func (e Entry) String() string {
	return fmt.Sprintf("%s %s, %d", e.Name, e.Surname, e.Year)
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

func zeroPtoS() *Entry {
	t := &Entry{}
	return t
}


func initPToS(N, S string, Y int) *Entry {
	if Y < 2000 {
		return &Entry{Name: N, Surname: S, Year: 2000}
	}

	return &Entry{Name: N, Surname: S, Year: Y}
}


func main() {
	x := zeroS()
	x.Name = "Leo"
	fmt.Println(x)
}
