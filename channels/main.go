package main

import (
	"fmt"
	"time"
)

// Function to demonstrate sending data through a channel
func sendData(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(500 * time.Millisecond)
	}
	close(ch)
}

// Function to demonstrate receiving data from a channel
func receiveData(ch chan int) {
	for {
		select {
		case data, ok := <-ch:
			if !ok {
				fmt.Println("Channel closed")
				return
			}
			fmt.Println("Received data:", data)
		}
	}
}

func main() {
	ch := make(chan int)

	go sendData(ch)
	go receiveData(ch)

	time.Sleep(3 * time.Second)
}
