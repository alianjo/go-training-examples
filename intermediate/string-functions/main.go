package main

import (
	"fmt"
	"strings"
)

var P = fmt.Println

func main() {
	P("Check character existence: ", strings.Contains("test", "t"))
	P("count :", strings.Count("alili", "li"))
	P("Join :", strings.Join([]string{"John", "Richard", "Alex", "Emma"}, ", "))
	P("Replace", strings.Replace("I love to eat pizza and pizza is my favorite restaurant", "pizza", "burger", 1))
	P("Split :", strings.Split("John,Alex,Emma", ","))
}
