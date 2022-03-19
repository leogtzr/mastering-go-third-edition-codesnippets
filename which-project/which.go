package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func listProgram(arg string) {
	file := arg
	path := os.Getenv("PATH")
	// a portable way of separating a list of paths:
	pathSplit := filepath.SplitList(path)
	for _, directory := range pathSplit {
		fullPath := filepath.Join(directory, file)
		// Does it exist?
		fileInfo, err := os.Stat(fullPath)
		if err == nil {
			if runtime.GOOS == "windows" {
				fmt.Println(fullPath)
			} else {
				mode := fileInfo.Mode()
				// Is it a regular file?
				if mode.IsRegular() {
					// Is it executable?
					if mode&0111 != 0 {
						fmt.Println(fullPath)
						// return
					}
				}
			}
		}
	}
}

func main() {

	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Fprintln(os.Stderr, "Please provide an argument")
		return
	}

	for _, arg := range arguments[1:] {
		listProgram(arg)
	}
}
