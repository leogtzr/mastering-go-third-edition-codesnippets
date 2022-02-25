package main

import (
	"fmt"
)

func main() {
	// Get User Input
	fmt.Printf("Name: ")
	var name string
	fmt.Scanln(&name)
	fmt.Printf("Your name is: %s\n", name)
}
