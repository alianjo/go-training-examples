package main

import (
	"fmt"
	"os"
)

func main() {
	FILE_PATH := "/tmp/mytestfile"
	f, err := os.OpenFile(FILE_PATH, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString("Hello, World!\n")
	f.WriteString("This is appended text.\n")
	fmt.Printf("file created at %s", FILE_PATH)
}
