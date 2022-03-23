package main

import (
	"fmt"
	"time"
)

func main() {
	loc, _ := time.LoadLocation("America/New_York")
	fmt.Printf("[%s]\n", time.Now().In(loc))
}
