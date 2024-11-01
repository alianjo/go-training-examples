package main

import (
	"fmt"
	"time"
)

func main() {
	P := fmt.Println
	location, _ := time.LoadLocation("Asia/Tehran")
	P(location)
	P(time.Now().In(location))
	then := time.Date(2009, 11, 17, 20, 34, 58, 53232323, time.UTC)
	P(then)
	P(then.Year())
	P(then.Weekday())
	P(then.Before(time.Now()))
	P(then.After(time.Now()))
	diff := time.Now().Sub(then)
	P(diff)
	P(diff.Hours())
	P(diff.Microseconds())

}
