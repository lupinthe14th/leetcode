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
