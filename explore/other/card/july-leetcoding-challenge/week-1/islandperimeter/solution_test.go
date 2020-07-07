package islandperimeter

import (
	"fmt"
	"testing"
)

func TestIslandPerimeter(t *testing.T) {
	tests := []struct {
		in   [][]int
		want int
	}{
		{in: [][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}, want: 16},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := islandPerimeter(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
