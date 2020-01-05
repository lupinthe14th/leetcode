package main

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
	arr     []int
	queries [][]int
}

var cases = []Case{
	{input: input{arr: []int{1, 3, 4, 8}, queries: [][]int{{0, 1}, {1, 2}, {0, 3}, {3, 3}}}, want: []int{2, 7, 14, 8}},
	{input: input{arr: []int{4, 8, 2, 10}, queries: [][]int{{2, 3}, {1, 3}, {0, 0}, {0, 3}}}, want: []int{8, 0, 4, 4}},
}

func TestXorQueries(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := xorQueries(tt.input.arr, tt.input.queries)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%v, want: %v", got, tt.want)
			}
		})
	}
}
