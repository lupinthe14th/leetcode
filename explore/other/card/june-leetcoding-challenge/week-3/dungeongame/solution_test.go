package dungeongame

import (
	"fmt"
	"testing"
)

func TestCalculateMinimumHP(t *testing.T) {
	tests := []struct {
		in   [][]int
		want int
	}{
		{in: [][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}}, want: 7},
		{in: [][]int{{100}}, want: 1},
		{in: [][]int{{-3, 5}}, want: 4},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := calculateMinimumHP(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
