package countnegativenumbersinasortedmatrix

import (
	"fmt"
	"testing"
)

type Case struct {
	in   [][]int
	want int
}

var cases = []Case{
	{in: [][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}}, want: 8},
	{in: [][]int{{3, 2}, {1, 0}}, want: 0},
	{in: [][]int{{1, -1}, {-1, -1}}, want: 3},
	{in: [][]int{{-1}}, want: 1},
}

func TestCountNegatives(t *testing.T) {
	for i, c := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := countNegatives(c.in)
			if got != c.want {
				t.Errorf("%d, want: %d", got, c.want)
			}
		})
	}
}
