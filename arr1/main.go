package main

import (
    "fmt"
)

func main() {
    x := [2]string{"Leonardo", "Gutierrez"}
    for i, v := range x {
        fmt.Println(i, v)
    }
}

