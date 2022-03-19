package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func main() {
	LOGFILE := path.Join(os.TempDir(), "mGo.log")
	file, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// The call to os.OpenFile() creates the log file for writing,
	// if it does not already exist, or opens it for writing
	// by appending new data at the end of it (os.O_APPEND)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	iLog := log.New(file, "iLog", log.LstdFlags)
	iLog.Println("Hello, world!")
	iLog.Println("Mastering Go 3rd edition")
}
