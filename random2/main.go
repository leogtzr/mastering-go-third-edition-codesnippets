package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	MIN = 33
	MAX = 93
)

func random(min, max int) int {
	return rand.Intn(max - min) + min 
}

func getString(len int64) string {
	temp := ""
	startChar := "!"

	var i int64
	i = 0
	for {
		myRand := random(MIN, MAX)
		newChar := string(startChar[0] + byte(myRand))
		temp += newChar
		if i == len {
			break
		}
		i++
	}

	return temp
}


func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(random(1, 100))
	fmt.Println(getString(20))
}
