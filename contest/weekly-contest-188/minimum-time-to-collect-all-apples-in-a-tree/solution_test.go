package minimumtimetocollectallapplesinatree

import (
	"fmt"
	"testing"
)

func TestMinTime(t *testing.T) {
	type in struct {
		n        int
		edges    [][]int
		hasApple []bool
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{n: 7, edges: [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}}, hasApple: []bool{false, false, true, false, true, true, false}}, want: 8},
		{in: in{n: 7, edges: [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}}, hasApple: []bool{false, false, true, false, false, true, false}}, want: 6},
		{in: in{n: 7, edges: [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}}, hasApple: []bool{false, false, false, false, false, false, false}}, want: 0},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := minTime(tt.in.n, tt.in.edges, tt.in.hasApple)
			if got != tt.want {
				t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
			}
		})
	}
}
