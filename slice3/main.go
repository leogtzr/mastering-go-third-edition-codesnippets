package main

import (
	"fmt"
)

func main() {
	a := make([]int, 4)
	fmt.Println("L:", len(a), "C:", cap(a))
	a = append(a, 5)
	fmt.Println("L:", len(a), "C:", cap(a))
}
