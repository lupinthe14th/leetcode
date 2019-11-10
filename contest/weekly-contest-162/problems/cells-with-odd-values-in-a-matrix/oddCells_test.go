package cellswithoddvaluesinamatrix

import (
	"fmt"
	"testing"
)

type input struct {
	n       int
	m       int
	indices [][]int
}

var cases = []struct {
	id    int
	input input
	want  int
}{
	{id: 1, input: input{n: 2, m: 3, indices: [][]int{{0, 1}, {1, 1}}}, want: 6},
	{id: 2, input: input{n: 2, m: 2, indices: [][]int{{1, 1}, {0, 0}}}, want: 0},
}

func TestOddCells(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := oddCells(tt.input.n, tt.input.m, tt.input.indices)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
