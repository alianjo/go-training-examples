package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := os.ReadFile("/tmp/dat")
	check(err)
	fmt.Print(string(dat))

	f, err := os.Open("/tmp/dat")
	check(err)
	b := make([]byte, 5)
	n1, err := f.Read(b)
	fmt.Printf("%d bytes : %s\n", n1, string(b[:n1]))

}
