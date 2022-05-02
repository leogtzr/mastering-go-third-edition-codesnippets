package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
)

const (
	MIN = 33
	MAX = 93
)

func generateBytes(n int64) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func main() {
	bytes, err := generateBytes(10)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(base64.URLEncoding.EncodeToString(bytes))
}
