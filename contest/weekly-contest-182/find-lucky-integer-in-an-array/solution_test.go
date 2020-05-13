package findluckyintegerinanarray

import (
	"fmt"
	"testing"
)

type Case struct {
	in   []int
	want int
}

var cases = []Case{
	{in: []int{2, 2, 3, 4}, want: 2},
	{in: []int{1, 2, 2, 3, 3, 3}, want: 3},
	{in: []int{2, 2, 2, 3, 3}, want: -1},
	{in: []int{5}, want: -1},
	{in: []int{7, 7, 7, 7, 7, 7, 7}, want: 7},
}

func TestFindLucky(t *testing.T) {
	for i, c := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := findLucky(c.in)
			if got != c.want {
				t.Errorf("got: %v, want: %v", got, c.want)
			}
		})
	}
}
