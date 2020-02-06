package thekweakestrowsinamatrix

import (
	"fmt"
	"reflect"
	"testing"
)

type Case struct {
	input input
	want  []int
}

type input struct {
	mat [][]int
	k   int
}

var cases = []Case{
	{input: input{mat: [][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 1},
	}, k: 3}, want: []int{2, 0, 3}},
}

func TestKWeakestRows(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := kWeakestRows(tt.input.mat, tt.input.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("input: %+v, got: %v, want: %v", tt.input, got, tt.want)
			}
		})
	}
}
