package searchinsertposition

import (
	"fmt"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	type in struct {
		nums   []int
		target int
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{nums: []int{1, 3, 5, 6}, target: 5}, want: 2},
		{in: in{nums: []int{1, 3, 5, 6}, target: 2}, want: 1},
		{in: in{nums: []int{1, 3, 5, 6}, target: 7}, want: 4},
		{in: in{nums: []int{1, 3, 5, 6}, target: 0}, want: 0},
		{in: in{nums: []int{1, 3}, target: 2}, want: 1},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := searchInsert(tt.in.nums, tt.in.target)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
