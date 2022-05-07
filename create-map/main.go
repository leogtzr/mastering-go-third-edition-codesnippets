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

	map2["maria"] = 51
	fmt.Println(map2)
	delete(map2, "maria")
	fmt.Println(map2)
	
	fmt.Println(len(map2))

	if v, ok := map2["leonardo"]; ok {
		fmt.Println(ok)
		fmt.Println(v)
	}
}
