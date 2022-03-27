package main

import "fmt"

func main() {
	// Create an empty slice
	aSlice := []float64{}
	// Both length and capacity are 0 because aSlice is empty
	fmt.Println(aSlice, len(aSlice), cap(aSlice))
	// Add elements to a slice
	aSlice = append(aSlice, 1234.56)
	aSlice = append(aSlice, -34.0)
	fmt.Println(aSlice, "with length", len(aSlice))
	fmt.Println(aSlice, len(aSlice), cap(aSlice))

	t := make([]int, 4)
	t[0] = 1
	t[1] = 2
	t[2] = 3
	t[3] = 4

	fmt.Printf("t -> %d, %d\n", len(t), cap(t))

	t = append(t, 5)
	fmt.Printf("t -> %d, %d\n", len(t), cap(t))
}
