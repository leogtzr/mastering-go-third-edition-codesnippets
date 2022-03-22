package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 23
	input := strconv.Itoa(n)
	fmt.Println(input)
	fmt.Printf("%T -> %v\n", input, input)

	input2 := strconv.FormatInt(int64(n), 10)
	fmt.Printf("%T -> %v\n", input2, input2)

	x := 54
	fmt.Printf("{%T}\n", x)

	var a int64 = 1
	b := 1
	fmt.Println(a == int64(b))
}
