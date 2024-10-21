package main

import (
	"reflect"
	"testing"
)

func TestObjectTaker(t *testing.T) {
	tests := []struct {
		name string
		args []any
		want any
	}{
		{"single argument", []any{1}, int},
		// {"multiple arguments", []any{1, 2, 3}, []any{1, 2, 3}},
		// {"no arguments", []any{}, []any{}},
		// {"mixed type arguments", []any{1, "hello", true}, []any{1, "hello", true}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ObjectTaker(tt.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ObjectTaker() = %v, want %v", got, tt.want)
			}
		})
	}
}
