package reverseinteger

import (
	"fmt"
	"testing"
)

var cases = []struct {
	input int
	want  int
}{
	{input: 123, want: 321},
	{input: -123, want: -321},
	{input: 120, want: 21},
	{input: 1, want: 1},
	{input: 900000, want: 9},
	{input: 1534236469, want: 0},
}

func TestReverse(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprintln(tt.input), func(t *testing.T) {
			got := reverse(tt.input)
			if got != tt.want {
				t.Errorf("%v, want %v", got, tt.want)
			}
		})
	}
}
