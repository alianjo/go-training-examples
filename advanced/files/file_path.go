package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	// Join multiple path components into a single path
	p := filepath.Join("/tmp", "mytestfile")
	fmt.Println(p)

	// Join multiple path components into a single path, handling duplicate separators
	fmt.Println(filepath.Join("dir1//", "filename"))

	// Get the directory part of a path
	fmt.Println("Dir(p):", filepath.Dir(p))

	// Get the base name of a path
	fmt.Println("Base(p):", filepath.Base(p))

	// Check if a path is absolute
	fmt.Println(filepath.IsAbs(p))

	// Get the file extension of a path
	myconfig := "myconfig.json"
	ext := filepath.Ext(myconfig)
	fmt.Println(ext)

	// Remove the file extension from a path
	fmt.Println(strings.TrimSuffix(myconfig, ext))
}
