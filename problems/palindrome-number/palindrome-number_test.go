package palindromenumber

import (
	"fmt"
	"testing"
)

var cases = []struct {
	input int
	want  bool
}{
	{input: 121, want: true},
	{input: 1221, want: true},
	{input: 12321, want: true},
	{input: -121, want: false},
	{input: 10, want: false},
}

func TestIsPalindrome(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprintln(tt.input), func(t *testing.T) {
			got := isPalindrome(tt.input)
			if got != tt.want {
				t.Errorf("%t, want %t", got, tt.want)
			}
		})
	}
}
