package main

import (
	"fmt"
	"io/ioutil"
	"sync"
)

// readFile reads a file and sends the result to a channel
func readFile(filename string, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		results <- fmt.Sprintf("Error reading %s: %v", filename, err)
		return
	}

	results <- fmt.Sprintf("Content of %s:\n%s", filename, string(data))
}

// showingResult reads results from the channel using `select`
func showingResult(results <-chan string, done chan<- bool) {
	for {
		select {
		case result, ok := <-results:
			if !ok {
				done <- true // Signal completion
				return
			}
			fmt.Println(result) // Print result as soon as it's available
		}
	}
}

func main() {
	files := []string{"file1.txt", "file2.txt"}
	var wg sync.WaitGroup
	results := make(chan string, len(files))
	done := make(chan bool)

	// Start concurrent result printing
	go showingResult(results, done)

	// Launch concurrent file reading
	for _, file := range files {
		wg.Add(1)
		go readFile(file, &wg, results)
	}

	wg.Wait()
	close(results) // Close the channel to signal `showingResult` to stop
	<-done         // Wait for showingResult to finish
}
