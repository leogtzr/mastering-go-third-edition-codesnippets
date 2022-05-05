package main

import (
	"fmt"
)

func con(a1, a2 [2]int) []int {
	return append(a1[:], a2[:]...)
}

func con2(a1, a2 [2]int) [4]int {
	var arr [len(a1) + len(a2)]int
	
	copy(arr[:], con(a1, a2))

	return arr
}


func main() {
	a1 := [2]int{2, 3}
	a2 := [2]int{4, 5}
	
	fmt.Println(con(a1, a2))
}
