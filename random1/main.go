package main

import (
	"fmt"
	"math/rand"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max - min) + min 
}


func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(random(1, 100))
}
