package main

import "testing"

func TestCheckKeyExistence(t *testing.T) {
	tests := []struct {
		name string
		data map[string]string
		key  string
		want bool
	}{
		{"key exists", map[string]string{"key": "value"}, "key", true},
		{"key does not exist", map[string]string{"key": "value"}, "non-existent-key", false},
		{"empty map", map[string]string{}, "key", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkKeyExistence(tt.data, tt.key); got != tt.want {
				t.Errorf("checkKeyExistence() = %v, want %v", got, tt.want)
			}
		})
	}
}
