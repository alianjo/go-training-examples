package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s := "This is string"

	h := sha256.New()

	h.Write([]byte(s))

	fmt.Printf("%x ", h.Sum(nil))
}
