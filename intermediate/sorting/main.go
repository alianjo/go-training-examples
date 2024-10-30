package main

import (
	"fmt"
	"slices"
)

func main() {
	// Sorting string in a slice
	Mys := []string{"John", "Frank", "Alex"}
	slices.Sort(Mys)
	fmt.Println(Mys)

	// Sorting ints in slice
	Mynumbers := []int{
		999, 2, 13, 4, 333, 12, 5, 1, 2, 3, 111, 13, 333,
	}
	slices.Sort(Mynumbers)
	if slices.IsSorted(Mynumbers) == false {
		fmt.Println("Mynumbers slice is not sorted!")
	} else {
		fmt.Println("The Mynumbers slice has been sorted:", Mynumbers)
	}
}
