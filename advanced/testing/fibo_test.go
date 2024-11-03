package main

import (
	"fmt"
	"testing"
)

func TestFibo(t *testing.T) {
	res := fibonacci(6)
	if res != 8 {
		t.Errorf("Expected 8 but got %d", res)
	}
}

func TestMultipleData(t *testing.T) {
	var tests = []struct {
		a    int
		want int
	}{
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{10, 55},
	}
	for _, tt := range tests {
		test := fmt.Sprintf("test number: %d, expected result: %d", tt.a, tt.want)
		t.Run(test, func(t *testing.T) {
			resp := fibonacci(tt.a)
			if resp != tt.want {
				t.Errorf("Got error for input %d, expected %d but got %d", tt.a, tt.want, resp)
			}
		})
	}
}
