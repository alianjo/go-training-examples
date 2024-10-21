package main

import "fmt"

func ObjectTaker(i ...any) any {
	return i
}

func main() {
	i := ObjectTaker([]int{1, 2, 3, 4}, map[int]string{1: "John", 2: "Frank"}, 3)
	fmt.Print(i)

	nextInt := func() func() int {
		i := 0
		return func() int {
			i++
			return i
		}
	}
	nextIntFunc := nextInt()
	fmt.Println(nextIntFunc())
	fmt.Println(nextIntFunc())
	fmt.Println(nextIntFunc())
	fmt.Println(nextIntFunc())

}
