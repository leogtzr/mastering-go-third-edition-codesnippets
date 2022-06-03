package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	MIN = 33
	MAX = 93
)

// CSVFILE resides in the home directory of the current user
const CSVFILE = "/Users/leogtzr/csv.data"

// Entry ...
type Entry struct {
	Name       string
	Surname    string
	Tel        string
	LastAccess string
}

func (entry Entry) String() string {
	return fmt.Sprintf("%s %s, phone: %s", entry.Name, entry.Surname, entry.Tel)
}

var data = []Entry{}
var index map[string]int

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

func matchTel(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`\d+$`)

	return re.Match(t)
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

func readCSVFile(filepath string) error {
	_, err := os.Stat(filepath)
	if err != nil {
		return err
	}

	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	// CSV file read all at once
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return err
	}

	for _, line := range lines {
		temp := Entry{
			Name:       line[0],
			Surname:    line[1],
			Tel:        line[2],
			LastAccess: line[3],
		}
		// Storing to global variable
		data = append(data, temp)
	}

	return nil
}

func createIndex() error {
	index = make(map[string]int)
	for i, k := range data {
		index[k.Tel] = i
	}

	return nil
}

// Initialized by the user - returns a pointer
// If it returns nil, there was an error
func initS(N, S, T string) *Entry {
	if T == "" || S == "" {
		return nil
	}

	lastAccess := strconv.FormatInt(time.Now().Unix(), 10)

	return &Entry{Name: N, Surname: S, Tel: T, LastAccess: lastAccess}
}

func saveCSVFile(filepath string) error {
	csvFile, err := os.Create(filepath)
	if err != nil {
		return err
	}

	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)
	for _, row := range data {
		temp := []string{row.Name, row.Surname, row.Tel, row.LastAccess}
		_ = csvWriter.Write(temp)
	}

	csvWriter.Flush()

	return nil
}

func deleteEntry(key string) error {
	i, ok := index[key]
	if !ok {
		return fmt.Errorf("%s cannot be found!", key)
	}

	data = append(data[:i], data[i+1:]...)
	// Update the index - key does not exist any more
	delete(index, key)

	err := saveCSVFile(CSVFILE)
	if err != nil {
		return err
	}

	return nil
}

func insert(pS *Entry) error {
	// If it already exists, do not add it
	_, ok := index[(*pS).Tel]
	if ok {
		return fmt.Errorf("%s already exists", pS.Tel)
	}

	data = append(data, *pS)
	// Update the index
	_ = createIndex()

	err := saveCSVFile(CSVFILE)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		exe := path.Base(arguments[0])
		fmt.Printf("usage: %s search|list <arguments>\n", exe)
		os.Exit(0)
	}

	// If the CSVFILE does not exist, create an empty one
	_, err := os.Stat(CSVFILE)

	if err != nil {
		// fmt.Println("Creating", CSVFILE)
		f, err := os.Create(CSVFILE)
		if err != nil {
			f.Close()
			fmt.Println(err)
			return
		}
		f.Close()
	}

	fileInfo, _ := os.Stat(CSVFILE)
	// Is it a regular file?
	mode := fileInfo.Mode()
	if !mode.IsRegular() {
		fmt.Println(CSVFILE, "not a regular file!")
		return
	}

	err = readCSVFile(CSVFILE)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = createIndex()
	if err != nil {
		fmt.Println("Cannot create index")
		return
	}

	switch arguments[1] {
	case "insert":
		if len(arguments) != 5 {
			fmt.Println("Usage: insert Name Surname Telephone")
			return
		}
		t := strings.ReplaceAll(arguments[4], "-", "")
		if !matchTel(t) {
			fmt.Println("Not a valid telephone number:", t)
			return
		}

		if temp := initS(arguments[2], arguments[3], t); temp != nil {
			err := insert(temp)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	case "search":
		if len(arguments) != 3 {
			fmt.Println("usage: search Surname")
			os.Exit(1)
		}

		t := strings.ReplaceAll(arguments[2], "-", "")
		if !matchTel(t) {
			fmt.Println("Not a valid telephone number")
			return
		}

		temp := search(t)
		if temp == nil {
			fmt.Println("Number not found:", t)
			return
		}

		fmt.Println(*temp)

	case "list":
		list()

	case "delete":
		if len(arguments) != 3 {
			fmt.Println("Usage: delete number")
			return
		}

		t := strings.ReplaceAll(arguments[2], "-", "")
		if !matchTel(t) {
			fmt.Println("Not a valid telephone number")
			return
		}

		if err := deleteEntry(t); err != nil {
			fmt.Println(err)
		}

	default:
		fmt.Println("Not a valid option")
	}
}
