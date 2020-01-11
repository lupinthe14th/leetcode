package main

import (
	"fmt"
	"reflect"
	"testing"
)

type input struct {
	mat [][]int
	K   int
}

type Case struct {
	input input
	want  [][]int
}

var cases = []Case{
	{input: input{mat: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, K: 1}, want: [][]int{{12, 21, 16}, {27, 45, 33}, {24, 39, 28}}},
	{input: input{mat: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, K: 2}, want: [][]int{{45, 45, 45}, {45, 45, 45}, {45, 45, 45}}},
}

func TestMatrixBlockSum(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := matrixBlockSum(tt.input.mat, tt.input.K)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%v, want: %v", got, tt.want)
			}
		})
	}
}
