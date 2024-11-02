package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	DIR := "/tmp/tmpdir"
	_, err := os.Stat(DIR + "/empty_file.txt")
	if os.IsNotExist(err) {
		f, err := os.Create(DIR + "/empty_file.txt")
		check(err)
		defer f.Close()
	} else {
		check(err)
	}
	check(err)

	defer os.Remove(DIR)

	f, err := os.Create(DIR + "/empty_file.txt")
	check(err)
	defer f.Close()

	err = os.MkdirAll(filepath.Join(DIR, "library"), 0755)
	check(err)

	err = os.Chdir(DIR)
	check(err)

	c, err := os.ReadDir(".")
	check(err)

	fmt.Println("Listing the directory content...")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	fmt.Println("See all the sub-directories...")
	err = filepath.WalkDir(DIR, visit)

}

func visit(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", path, d.IsDir())
	return nil
}
