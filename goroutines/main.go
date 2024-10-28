package main

import (
	"fmt"
	"time"
)

var message = make(chan string, 3)

func Manager() {
	time.Sleep(time.Second)
	message <- "PING" // Sends the string "PING" to the message channel

}

func main() {

	go func() { fmt.Println("Start Goroutine function"); Manager() }()

	fmt.Println(<-message)

}
