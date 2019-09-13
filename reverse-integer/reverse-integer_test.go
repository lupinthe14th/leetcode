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
}

func TestReverse(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprintln(tt.input), func(t *testing.T) {
			actual := reverse(tt.input)
			if actual != tt.want {
				t.Errorf("%v, want %v", actual, tt.want)
			}
		})
	}
}
