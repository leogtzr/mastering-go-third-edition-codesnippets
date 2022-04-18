package main

import "fmt"

const (
	Zero = 0
	One = 1
	Two = 2
)


type Power2 int


func main() {
	fmt.Println(One)
    fmt.Println(Two)
    const (
        p2_0 Power2 = 1 << iota
        _
        p2_2
        _
        p2_4
        _
        p2_6
    )
    fmt.Println("2^0:", p2_0)
    fmt.Println("2^2:", p2_2)
    fmt.Println("2^4:", p2_4)
    fmt.Println("2^6:", p2_6)
}
