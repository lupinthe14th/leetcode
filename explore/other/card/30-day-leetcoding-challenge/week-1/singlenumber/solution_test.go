package singlenumber

import (
	"fmt"
	"testing"
)

type Case struct {
	in   []int
	want int
}

var cases = []Case{
	{in: []int{2, 2, 1}, want: 1},
	{in: []int{4, 1, 2, 1, 2}, want: 4},
}

func TestSingleNumber(t *testing.T) {
	for i, c := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := singleNumber(c.in)
			if got != c.want {
				t.Fatalf("got: %v, want: %v", got, c.want)
			}
		})
	}
}
