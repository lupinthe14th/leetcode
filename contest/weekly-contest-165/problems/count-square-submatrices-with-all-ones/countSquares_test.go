package countsquaresubmatriceswithallones

import (
	"fmt"
	"testing"
)

var cases = []struct {
	id    int
	input [][]int
	want  int
}{
	{id: 1, input: [][]int{{0, 1, 1, 1}, {1, 1, 1, 1}, {0, 1, 1, 1}}, want: 15},
	{id: 2, input: [][]int{{1, 0, 1}, {1, 1, 0}, {1, 1, 0}}, want: 7},
}

func TestCountSquares(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := countSquares(tt.input)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
