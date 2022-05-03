package main

import (
	"fmt"
	"math/rand"
	"os"
	"path"
	"strconv"
	"time"
)

const (
	MIN = 33
	MAX = 93
)

// Entry ...
type Entry struct {
	Name    string
	Surname string
	Tel     string
}

func (entry Entry) String() string {
	return fmt.Sprintf("%s %s, phone: %s", entry.Name, entry.Surname, entry.Tel)
}

var data = []Entry{}

func search(key string) *Entry {
	for i, v := range data {
		/*
			if v.Surname == key {
				return &data[i]
			}
		*/
		if v.Tel == key {
			fmt.Println("       FOUND")
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

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func getString(len int64) string {
	temp := ""
	startChar := "!"
	var i int64 = 1
	for {
		myRand := random(MIN, MAX)
		newChar := string(startChar[0] + byte(myRand))
		temp = temp + newChar
		if i == len {
			break
		}
		i++
	}
	return temp
}

func populate(n int) {
	for i := 0; i < n; i++ {
		name := getString(4)
		surname := getString(5)
		n := strconv.Itoa(random(100, 199))
		data = append(data, Entry{name, surname, n})
	}
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		exe := path.Base(arguments[0])
		fmt.Printf("usage: %s search|list <arguments>\n", exe)
		os.Exit(0)
	}

	rand.Seed(time.Now().UnixNano())

	populate(10)
	data = append(data, Entry{"Leo", "Gtz", "123"})

	fmt.Printf("%d entries\n", len(data))

	for _, d := range data {
		fmt.Println(d)
	}

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
