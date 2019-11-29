package lengthoflastword

import (
	"fmt"
	"testing"
)

var cases = []struct {
	id    int
	input string
	want  int
}{
	{id: 1, input: "Hello World", want: 5},
	{id: 2, input: "", want: 0},
	{id: 3, input: "Hello", want: 5},
	{id: 4, input: "Hello Golang World", want: 5},
	{id: 5, input: "No Silver Bullet in programming", want: 11},
	{id: 6, input: "a", want: 1},
}

func TestLengthOfLastWord(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := lengthOfLastWord(tt.input)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
