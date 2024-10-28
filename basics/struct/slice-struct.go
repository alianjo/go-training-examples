package main

import (
	"fmt"
)

func main() {
	students := []struct {
		Name        string
		Proficiency string // e.g., "Native", "Fluent", Intermediate"
	}{
		{"Ali", "Intermediate"},
		{"John", "Fluent"},
	}
	fmt.Print(students)
}
