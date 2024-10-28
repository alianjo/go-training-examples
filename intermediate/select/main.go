package main

import (
	"fmt"
	"time"
)

func waiter(ch chan<- bool) {
	time.Sleep(time.Second * 5)
	ch <- true
}

func main() {
	ch := make(chan bool, 1)
	go waiter(ch)
	select {
	case <-ch:
		fmt.Println("We got the ch")
	case <-time.After(time.Second * 10):
		fmt.Println("Timeout")
	}
}
