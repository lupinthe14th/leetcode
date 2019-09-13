package main

import (
	"fmt"
	"testing"
)

var cases = []struct {
	input int
	want  bool
}{
	{input: 121, want: true},
	{input: -121, want: false},
	{input: 10, want: false},
}

func TestIsPalindrome(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprintln(tt.input), func(t *testing.T) {
			actual := isPalindrome(tt.input)
			if actual != tt.want {
				t.Errorf("%t, want %t", actual, tt.want)
			}
		})
	}
}
