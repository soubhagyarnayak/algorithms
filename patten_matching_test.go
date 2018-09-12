package datastructure

import (
	"testing"
)

func TestFindSubstring(t *testing.T) {
	var tests = []struct {
		input_string    string
		input_substring string
		want            int
	}{
		{"banana", "nana", 2},
		{"abrakadabrakaabra", "dabra", 6},
		{"soubhagya r nayak", "nayak", 12},
		{"banana", "nanan", -1},
		{"this is a string with some characters", "this is a string with some special characters", -1},
		{"this time substring random will come in middle", "random", 20},
		{"beginning with the substring", "beginning", 0},
	}
	for _, test := range tests {
		if got := FindSubstring(test.input_string, test.input_substring); got != test.want {
			t.Errorf("Expected FindSubstring(%q,%q)=%v, but got %v", test.input_string, test.input_substring, test.want, got)
		}
	}
}
