package main

import "fmt"

func main() {
	var a [5]int

	fmt.Println(a)

	b := []int{1, 2, 3, 4, 5}

	fmt.Println(b)

	c := []int{100, 3: 400, 500}
	fmt.Println(c)

	f := [2][3]int{{5, 22, 11}, {4, 6}}
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(f[i][j])
		}
	}
	copy(b, c)
	fmt.Println(b)
}
