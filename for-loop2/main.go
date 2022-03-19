package main

import "fmt"

func main() {
	i := 0
	for ok := true; ok; ok = (i != 10) {
		fmt.Println(i*i, " ")
		i++
	}
	fmt.Println()
}


