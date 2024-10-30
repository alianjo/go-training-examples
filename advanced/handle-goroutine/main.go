package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i <= 5; i++ {
		wg.Add(1)
		time.Sleep(time.Second)
		go func() {
			defer wg.Done()
			fmt.Println("HIIIII Routine Number %d", i)
			time.Sleep(time.Second)
			fmt.Println("Finishe Routine Number %d", i)
		}()
	}
	wg.Wait()
}
