package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	ListDirCmd := flag.NewFlagSet("dir", flag.ExitOnError)
	dir_path := ListDirCmd.String("path", "", "Path to list all the files")

	File_list_Cmd := flag.NewFlagSet("file", flag.ExitOnError)
	file_path := File_list_Cmd.String("path", "", "Path to list all the files")

	if len(os.Args) < 2 {
		fmt.Println("expected 'dir' or 'file' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {

	case "dir":
		ListDirCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'dir':")
		files, err := os.ReadDir(*dir_path)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			for _, file := range files {
				if file.IsDir() {
					fmt.Println(file.Name())
				}
			}
		}
	case "file":
		File_list_Cmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'file':")
		files, err := os.ReadDir(*file_path)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			for _, file := range files {
				if !file.IsDir() {
					fmt.Println(file.Name())
				}
			}
		}
	default:
		fmt.Println("expected 'file' or 'dir' subcommands")
		os.Exit(1)
	}
}
