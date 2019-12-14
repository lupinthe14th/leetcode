package removecoveredintervals

import (
	"fmt"
	"testing"
)

var cases = []struct {
	id    int
	input [][]int
	want  int
}{
	{id: 1, input: [][]int{{1, 4}, {3, 6}, {2, 8}}, want: 2},
	{id: 2, input: [][]int{{1, 4}}, want: 1},
	{id: 3, input: [][]int{{1, 4}, {3, 6}}, want: 2},
	{id: 4, input: [][]int{{1, 4}, {2, 3}}, want: 1},
	{id: 5, input: [][]int{{34335, 39239}, {15875, 91969}, {29673, 66453}, {53548, 69161}, {40618, 93111}}, want: 2},
	{id: 6, input: [][]int{{66672, 75156}, {59890, 65654}, {92950, 95965}, {9103, 31953}, {54869, 69855}, {33272, 92693}, {52631, 65356}, {43332, 89722}, {4218, 57729}, {20993, 92876}}, want: 3},
}

func TestRemoveCoveredIntervals(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := removeCoveredIntervals(tt.input)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
