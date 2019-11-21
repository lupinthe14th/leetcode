package binarysearch

import (
	"fmt"
	"testing"
)

type input struct {
	nums   []int
	target int
}

var cases = []struct {
	id    int
	input input
	want  int
}{
	{id: 1, input: input{nums: []int{-1, 0, 3, 5, 9, 12}, target: 9}, want: 4},
	{id: 2, input: input{nums: []int{-1, 0, 3, 5, 9, 12}, target: 2}, want: -1},
	{id: 3, input: input{nums: []int{-1, 0, 3, 5, 9, 12}, target: 13}, want: -1},
}

func TestSearch(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := search(tt.input.nums, tt.input.target)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
