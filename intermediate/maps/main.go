package main

import "fmt"

func checkKeyExistence(data map[string]string, key string) bool {
	_, ok := data[key]
	return ok
}

func main() {
	StudentData := make(map[string]string)

	StudentData["name"] = "John"
	StudentData["Age"] = "22"

	fmt.Println(StudentData)

	delete(StudentData, "Age")

	fmt.Println(StudentData)

	a := checkKeyExistence(StudentData, "name")
	fmt.Println(a)
}
