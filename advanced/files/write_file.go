package main

import "os"

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	data := []byte("HI\nThis is John")
	err := os.WriteFile("/tmp/mytestfile", data, 0644)
	handleError(err)

}
