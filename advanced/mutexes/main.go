package main

import (
	"fmt"
	"sync"
)

type Stuck struct {
	mu       sync.Mutex
	products map[string]int
}

func (s *Stuck) sell(name string) {

	s.mu.Lock()
	// Locking is necessary to prevent concurrent map writes.
	defer s.mu.Unlock()
	// Unlocking is deferred to ensure the lock is released after the function call.
	s.products[name]--
}

func main() {
	s := Stuck{

		products: map[string]int{"pen": 10000, "laptop": 10000},
	}

	var wg sync.WaitGroup

	doSell := func(name string, n int) {
		for i := 0; i < n; i++ {
			s.sell(name)
		}
		wg.Done()
	}

	wg.Add(2)
	go doSell("pen", 10000)
	go doSell("laptop", 10000)

	wg.Wait()
	fmt.Println(s.products)
}
