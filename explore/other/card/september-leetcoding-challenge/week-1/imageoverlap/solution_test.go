package imageoverlap

import (
	"fmt"
	"testing"
)

func TestLargestOverlap(t *testing.T) {
	t.Parallel()
	type in struct {
		A, B [][]int
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{A: [][]int{{1, 1, 0}, {0, 1, 0}, {0, 1, 0}}, B: [][]int{{0, 0, 0}, {0, 1, 1}, {0, 0, 1}}}, want: 3},
		{in: in{
			A: [][]int{{0, 0, 0, 0, 0}, {0, 1, 0, 0, 0}, {0, 0, 1, 0, 0}, {0, 0, 0, 0, 1}, {0, 1, 0, 0, 1}},
			B: [][]int{{1, 0, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {0, 1, 1, 1, 1}, {1, 0, 1, 1, 1}}}, want: 5},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := largestOverlap(tt.in.A, tt.in.B)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
