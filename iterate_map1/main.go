package main

import (
	"fmt"
)

func main() {
	aMap := map[string]int{
		"leo": 31,
		"maria": 51,
	}
	
	aMap["edgar"] = 33

	for k, v := range aMap {
		fmt.Printf("%s - %d\n", k, v)
	}
}
