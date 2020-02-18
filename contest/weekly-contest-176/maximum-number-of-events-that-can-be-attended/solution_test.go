package maximumnumberofeventsthatcanbeattended

import (
	"fmt"
	"testing"
)

type Case struct {
	in   [][]int
	want int
}

var cases = []Case{
	{in: [][]int{{1, 2}, {2, 3}, {3, 4}}, want: 3},
	{in: [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 2}}, want: 4},
	{in: [][]int{{1, 4}, {4, 4}, {2, 2}, {3, 4}, {1, 1}}, want: 4},
	{in: [][]int{{1, 100000}}, want: 1},
	{in: [][]int{{1, 1}, {1, 2}, {1, 3}, {1, 4}, {1, 5}, {1, 6}, {1, 7}}, want: 7},
}

func TestMaxEvents(t *testing.T) {
	for i, c := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := maxEvents(c.in)
			if got != c.want {
				t.Errorf("%d, want: %d", got, c.want)
			}
		})
	}
}
