package countsquaresubmatriceswithallones

import (
	"fmt"
	"testing"
)

func TestCountSquares(t *testing.T) {
	tests := []struct {
		in   [][]int
		want int
	}{
		{in: [][]int{
			{0, 1, 1, 1},
			{1, 1, 1, 1},
			{0, 1, 1, 1},
		}, want: 15},
		{in: [][]int{
			{1, 0, 1},
			{1, 1, 0},
			{1, 1, 0},
		}, want: 7},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := countSquares(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
