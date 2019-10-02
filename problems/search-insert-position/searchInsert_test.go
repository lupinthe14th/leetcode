package searchinsertposition

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
	{id: 1, input: input{nums: []int{1, 3, 5, 6}, target: 5}, want: 2},
	{id: 2, input: input{nums: []int{1, 3, 5, 6}, target: 2}, want: 1},
	{id: 3, input: input{nums: []int{1, 3, 5, 6}, target: 7}, want: 4},
	{id: 4, input: input{nums: []int{1, 3, 5, 6}, target: 0}, want: 0},
	{id: 5, input: input{nums: []int{1, 3, 5, 6}, target: -1}, want: 0},
}

func TestSearchInsert(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := searchInsert(tt.input.nums, tt.input.target)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
