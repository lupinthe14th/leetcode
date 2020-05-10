package counttripletsthatcanformtwoarraysofequalxor

import (
	"fmt"
	"testing"
)

func TestCountTriplets(t *testing.T) {
	tests := []struct {
		in   []int
		want int
	}{
		{in: []int{2, 3, 1, 6, 7}, want: 4},
		{in: []int{1, 1, 1, 1, 1}, want: 10},
		{in: []int{2, 3}, want: 0},
		{in: []int{1, 3, 5, 7, 9}, want: 3},
		{in: []int{7, 11, 12, 9, 5, 2, 7, 17, 22}, want: 8},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := countTriplets(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
			}
		})
	}
}
