package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	p := fmt.Println

	p(rand.IntN(100))

	p(rand.Float64())

	s3 := rand.NewPCG(900, 1024)
	r3 := rand.New(s3)
	fmt.Print(r3.IntN(100))
	fmt.Print(r3.IntN(100))
	fmt.Println()
}
