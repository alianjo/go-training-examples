package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	fruits := []string{"peach", "banana", "kiwi"}

	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}
	// This is a length-based sort, sorting the fruits slice in ascending order based on the length of each fruit's name.
	slices.SortFunc(fruits, lenCmp)
	fmt.Println(fruits)

	type Person struct {
		name string
		age  int
	}

	people := []Person{
		Person{name: "Jax", age: 37},
		Person{name: "TJ", age: 25},
		Person{name: "Alex", age: 72},
		Person{name: "Alex2", age: 43},
		Person{name: "John", age: 48},
	}
	// Age sorting
	slices.SortFunc(people,
		func(a, b Person) int {
			return cmp.Compare(a.age, b.age)
		})
	fmt.Println(people)

	type Students struct {
		Name string
		Id   int64
	}
	// Alphabet sorting
	cambridge_students := []Students{
		{Name: "Ali", Id: 4},
		{Name: "Ahmad", Id: 3},
		{Name: "AAA", Id: 5},
	}
	slices.SortFunc(cambridge_students,
		func(a, b Students) int {
			return cmp.Compare(a.Name, b.Name)
		},
	)
	fmt.Println(cambridge_students)

}
