package searchinrotatedsortedarray

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	type in struct {
		nums   []int
		target int
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 0}, want: 4},
		{in: in{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 3}, want: -1},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := search(tt.in.nums, tt.in.target)
			if got != tt.want {
				t.Fatalf("got: %v, want: %v", got, tt.want)
			}
		})
	}

}
