package besttimetobuyandsellstockwithcooldown

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	t.Parallel()
	tests := []struct {
		in   []int
		want int
	}{
		{in: []int{1, 2, 3, 0, 2}, want: 3},
		{in: []int{1, 2, 4}, want: 3},
		{in: []int{}, want: 0},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := maxProfit(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
