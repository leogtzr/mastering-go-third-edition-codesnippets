package main

import (
    "fmt"
)

type (
    Digit int
    Power2 int
)

const PI = 3.1415926

const (
    C1 = "C1C1C1"
    C2 = "C2C2C2"
    C3 = "C3C3C3"
)

func main() {
    fmt.Println(C1)
    fmt.Println(PI)
    const s1 = 123
    var v1 float32 = s1 * 12
    fmt.Println(v1)
    fmt.Println(PI)
    const (
        Zero Digit = iota
        One
        Two
        Three
        Four
    )
    fmt.Println(Zero)
    fmt.Println(One)
    fmt.Println(Two)
    fmt.Println(Three)
    fmt.Println(Four)
}
