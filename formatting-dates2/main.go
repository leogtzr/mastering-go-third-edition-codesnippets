package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	arg := os.Args[1]
	mydate, err := time.Parse("2006-01-02", arg)
	if err != nil {
		panic(err)
	}

	fmt.Println(mydate)
	fmt.Println(mydate.Day())
	fmt.Println(mydate.Month())
	fmt.Println(mydate.Year())
}
