package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	args := []string{"lsss", "-aa", "-l", "-h"}

	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
