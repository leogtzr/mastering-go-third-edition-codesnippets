package main

import "fmt"

func main() {
	a := make([]int, 3)
	fmt.Println(a)

	a = append(a, []int{4, 5, 6}...)
	fmt.Println(a)
}
