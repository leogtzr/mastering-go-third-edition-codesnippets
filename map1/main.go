package main

import (
	"fmt"
)

func main() {
	aMap := map[string]int{}
	fmt.Println(aMap)
	aMap = nil
	if (aMap == nil) {
		fmt.Println("es nil")
		aMap = map[string]int{
			"leo": 31,
		}
	}
	fmt.Println(aMap)
}
