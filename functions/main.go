package main

import "fmt"

func ObjectTaker(i ...any) any {
	return i
}

func main() {
	i := ObjectTaker([]int{1, 2, 3, 4}, map[int]string{1: "ali", 2: "sadeq"}, 3)
	fmt.Print(i.(int))
}
