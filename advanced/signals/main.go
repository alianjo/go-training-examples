package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan bool, 1)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println("Recieved sig:", sig)
		done <- true
	}()

	fmt.Println("Awaiting signal")
	<-done
	fmt.Println("Finish")
}
