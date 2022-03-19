package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	invalids := make([]string, 0)
	for _, k := range arguments[1:] {
		_, err := strconv.Atoi(k)
		if err != nil {
			invalids = append(invalids, k)
		}
	}

	fmt.Println(invalids)
}
