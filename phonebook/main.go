package main

import (
	"fmt"
	"os"
	"path"
)

// Entry ...
type Entry struct {
	Name    string
	Surname string
	Tel     string
}

var data = []Entry{}

func search(key string) *Entry {
	for i, v := range data {
		if v.Surname == key {
			return &data[i]
		}
	}

	return nil
}

func list() {
	for _, v := range data {
		fmt.Println(v)
	}
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		exe := path.Base(arguments[0])
		fmt.Printf("usage: %s search|list <arguments>\n", exe)
		os.Exit(0)
	}

	data = append(data, Entry{"Mihalis", "Tsoukalos", "2109416471"})
	data = append(data, Entry{"Mary", "Doe", "2109416871"})
	data = append(data, Entry{"John", "Black", "2109416123"})

	switch arguments[1] {
	case "search":
		if len(arguments) != 3 {
			fmt.Println("usage: search Surname")
			os.Exit(1)
		}

		result := search(arguments[2])
		if result == nil {
			fmt.Println("Entry not found:", arguments[2])
			return
		}

		fmt.Println(*result)

	case "list":
		list()

	default:
		fmt.Println("Not a valid option")
	}
}