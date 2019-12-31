package main

import (
	"fmt"
	"testing"
)

type input struct {
	arr   []int
	start int
}

var cases = []struct {
	id    int
	input input
	want  bool
}{
	{id: 1, input: input{arr: []int{4, 2, 3, 0, 3, 1, 2}, start: 5}, want: true},
	{id: 2, input: input{arr: []int{4, 2, 3, 0, 3, 1, 2}, start: 0}, want: true},
	{id: 3, input: input{arr: []int{3, 0, 2, 1, 2}, start: 2}, want: false},
}

func TestCanReach(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := canReach(tt.input.arr, tt.input.start)
			if got != tt.want {
				t.Errorf("%t, want: %t", got, tt.want)
			}
		})
	}
}
