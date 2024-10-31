package main

import (
	"errors"
	"fmt"
)

func division(i int) (result int, err error) {
	// Caching the zero division error
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in division:", r)
			err = errors.New(fmt.Sprintf("%v", r))
		}
	}()

	result = 200 / i
	return
}

func main() {
	result, err := division(0)
	if err != nil {
		fmt.Println("Error in main:", err)
	} else {
		fmt.Println("Result dividing 200 / 0 :", result)
	}

	result, err = division(5)
	if err != nil {
		fmt.Println("Error in main:", err)
	} else {
		fmt.Println("Result dividing 200 / 5 :", result)
	}
}
