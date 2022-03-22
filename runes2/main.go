package main

import "fmt"

func main() {
	r := '€'
	fmt.Println(r)
	fmt.Printf("c -> %c\n", r)

	for x, y := range string(r) {
		fmt.Println(x, y)
	}

	aString := "Hello World! €"
	fmt.Println("First character", string(aString[0]))

	fmt.Println("As an int32 value:", r)
	// Convert Runes to text
	fmt.Printf("As a string: %s and as a character: %c\n", r, r)
	// Print an existing string as runes
	for _, v := range aString {
		fmt.Printf("%x ", v)
	}
	fmt.Println()

	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	for _, v := range aString {
		fmt.Printf("%c", v)
	}
	fmt.Println()
}
