package main

import (
	"fmt"
)

func main() {
	x := 124
	y := &x

	fmt.Println(y)
	fmt.Println(*y)
}

