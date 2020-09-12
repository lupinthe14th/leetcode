package maximumproductsubarray

import (
	"fmt"
	"testing"
)

func TestMaxProduct(t *testing.T) {
	t.Parallel()
	tests := []struct {
		in   []int
		want int
	}{
		{in: []int{2, 3, -2, 4}, want: 6},
		{in: []int{-2, 0, -1}, want: 0},
		{in: []int{-2}, want: -2},
		{in: []int{0, 2}, want: 2},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := maxProduct(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
