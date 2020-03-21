package findthedistancevaluebetweentwoarrays

import (
	"fmt"
	"testing"
)

type in struct {
	arr1, arr2 []int
	d          int
}

type Case struct {
	in   in
	want int
}

var cases = []Case{
	{in: in{arr1: []int{4, 5, 8}, arr2: []int{10, 9, 1, 8}, d: 2}, want: 2},
	{in: in{arr1: []int{1, 4, 2, 3}, arr2: []int{-4, -3, 6, 10, 20, 30}, d: 3}, want: 2},
	{in: in{arr1: []int{2, 1, 100, 3}, arr2: []int{-5, -2, 10, -3, 7}, d: 6}, want: 1},
}

func TestFindTheDistanceValue(t *testing.T) {
	for i, c := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := findTheDistanceValue(c.in.arr1, c.in.arr2, c.in.d)
			if got != c.want {
				t.Errorf("got: %d, want: %d", got, c.want)
			}
		})
	}
}
