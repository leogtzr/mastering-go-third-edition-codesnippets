package main

import "fmt"

func change(s []string) {
	s[0] = "Change_function"
}

func main() {
	a := [4]string{"Zero", "One", "Two", "Three"}
	fmt.Println("a:", a)

	var S0 = a[0:1]
	fmt.Println(S0)

	S0[0] = "Cero"
	fmt.Println("a:", a)

	change(S0)

	fmt.Println("a:", a)

	S0 = append(S0, "Leonardo")
	S0 = append(S0, "Gtz")
	fmt.Println("a:", a)
	fmt.Println("S0:", S0)

}
