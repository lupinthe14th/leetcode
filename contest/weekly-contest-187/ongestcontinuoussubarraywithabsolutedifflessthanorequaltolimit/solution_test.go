package ongestcontinuoussubarraywithabsolutedifflessthanorequaltolimit

import (
	"fmt"
	"testing"
)

func TestLongestSubarray(t *testing.T) {
	type in struct {
		nums  []int
		limit int
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{nums: []int{8, 2, 4, 7}, limit: 4}, want: 2},
		{in: in{nums: []int{10, 1, 2, 4, 7, 2}, limit: 5}, want: 4},
		{in: in{nums: []int{4, 2, 2, 2, 4, 4, 2, 2}, limit: 0}, want: 3},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := longestSubarray(tt.in.nums, tt.in.limit)
			if got != tt.want {
				t.Fatalf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
