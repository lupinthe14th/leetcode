package main

import (
	"fmt"
	"testing"
)

type input struct {
	arr    []int
	target int
}

var cases = []struct {
	id    int
	input input
	want  int
}{
	{id: 1, input: input{arr: []int{4, 9, 3}, target: 10}, want: 3},
	{id: 2, input: input{arr: []int{2, 3, 5}, target: 10}, want: 5},
	{id: 3, input: input{arr: []int{60864, 25176, 27249, 21296, 20204}, target: 56803}, want: 11361},
	{id: 4, input: input{arr: []int{1547, 83230, 57084, 93444, 70879}, target: 71237}, want: 17422},
}

func TestFindBestValue(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := findBestValue(tt.input.arr, tt.input.target)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
