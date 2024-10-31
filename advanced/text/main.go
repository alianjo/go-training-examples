package main

import (
	"bytes"
	"fmt"
	"regexp"
)

var P = fmt.Println

func main() {
	r, err := regexp.Compile("\\d{3}-\\d{3}-\\d{4}")
	if err != nil {
		P(err)
		return
	}
	example_string := "Phone number 123-456-7489"
	P(r.FindString(example_string))
	P(r.FindStringIndex(example_string))
	P(r.FindStringSubmatchIndex(example_string))
	// Feed function
	P(string(r.ReplaceAllFunc([]byte("XXX-XXX-XXXX"), bytes.ToUpper)))
}
