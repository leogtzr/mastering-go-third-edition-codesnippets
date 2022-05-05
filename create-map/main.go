package main

import (
	"fmt"
)

func main() {
	map1 := make(map[string]int)
	map1["leonardo"] = 31
	map1["brenda"] = 26

	map2 := map[string]int{
		"leonardo": 31,
		"brenda": 26,
	}

	fmt.Println(map1)
	fmt.Println(map2)
}
