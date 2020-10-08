package removecoveredintervals

import (
	"fmt"
	"testing"
)

func TestRemoveCoveredIntervals(t *testing.T) {
	t.Parallel()
	type in struct {
		intervals [][]int
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{intervals: [][]int{{1, 4}, {3, 6}, {2, 8}}}, want: 2},
		{in: in{intervals: [][]int{{1, 4}, {2, 3}}}, want: 1},
		{in: in{intervals: [][]int{{0, 10}, {5, 12}}}, want: 2},
		{in: in{intervals: [][]int{{3, 10}, {4, 10}, {5, 11}}}, want: 2},
		{in: in{intervals: [][]int{{1, 2}, {1, 4}, {3, 4}}}, want: 1},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := removeCoveredIntervals(tt.in.intervals)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
