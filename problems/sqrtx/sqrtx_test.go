package sqrtx

import (
	"fmt"
	"testing"
)

var cases = []struct {
	id int

	input int
	want  int
}{
	{id: 1, input: 4, want: 2},
	{id: 2, input: 8, want: 2},
	{id: 3, input: 9, want: 3},
	{id: 4, input: 10, want: 3},
	{id: 5, input: 3, want: 1},
	{id: 6, input: 2, want: 1},
	{id: 7, input: 1, want: 1},
	{id: 8, input: 0, want: 0},
}

func TestMySqrt(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := mySqrt(tt.input)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
