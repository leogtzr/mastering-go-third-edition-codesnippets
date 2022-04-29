package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func (p *Person) String() string {
	return fmt.Sprintf("'%s' -> %d", p.Name, p.Age)
}

func main() {
	var leo *Person
	
	if leo == nil {
		fmt.Println("it is nil ... ")
		leo = new(Person)
		leo.Name = "Leonardo"
		(*leo).Name = "Leito"
	}

	fmt.Println(leo)
	fmt.Printf("Person -> [%+v]]\n", leo)
	fmt.Println(&leo)
}
	
