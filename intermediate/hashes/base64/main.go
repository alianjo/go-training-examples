package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	string_data := "asdfsdfsdfewr324342;ljj;v32!@#$%#"

	string_base64 := base64.StdEncoding.EncodeToString([]byte(string_data))
	fmt.Println(string_base64)

	string_decode_base64, _ := base64.StdEncoding.DecodeString(string_base64)
	if string_data == string(string_decode_base64) {
		fmt.Println(string(string_decode_base64))
		fmt.Println("Successfully decoded base64")
	}

}
