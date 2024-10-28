package main

import (
	"fmt"
)

func main() {
	// case statement
	TypeDefine := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("It's a boolean")
		case int:
			fmt.Println("This is an integer")
		case int8:
			fmt.Println("This is an int8")
		case int16:
			fmt.Println("This is an int16")
		case int32:
			fmt.Println("This is an int32")
		case int64:
			fmt.Println("This is an int64")
		case uint:
			fmt.Println("This is an uint")
		case uint8:
			fmt.Println("This is an uint8")
		case uint16:
			fmt.Println("This is an uint16")
		case uint32:
			fmt.Println("This is an uint32")
		case uint64:
			fmt.Println("This is an uint64")
		case uintptr:
			fmt.Println("This is an uintptr")
		case float32:
			fmt.Println("This is a float32")
		case float64:
			fmt.Println("This is a float64")
		case complex64:
			fmt.Println("This is a complex64")
		case complex128:
			fmt.Println("This is a complex128")
		case string:
			fmt.Println("This is a string")
		case error:
			fmt.Println("This is an error")
		case chan byte:
			fmt.Println("This is a channel")
		case map[string]int:
			fmt.Println("This is a map")
		case chan int:
			fmt.Println("this is a intiger channel")
		case struct{}:
			fmt.Println("This is a struct")
		case []string:
			fmt.Println("This is a slice")
		case []int:
			fmt.Println("This is an array")
		case *int:
			fmt.Println("This is a pointer")
		case *string:
			fmt.Println("This is a string pointer")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	TypeDefine(true)
	TypeDefine(10)
	TypeDefine(10.5)
	TypeDefine("Hello")
	TypeDefine(struct{}{})
	TypeDefine([]string{"a", "b", "c"})
	TypeDefine([]int{1, 2, 3})
	TypeDefine(map[string]int{"a": 1, "b": 2})
	TypeDefine(make(chan int))
	TypeDefine(complex(1, 2))
	TypeDefine(uintptr(10))

}
