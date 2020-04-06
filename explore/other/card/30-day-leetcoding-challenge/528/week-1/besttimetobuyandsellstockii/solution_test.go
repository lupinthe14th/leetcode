package besttimetobuyandsellstockii

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		in   []int
		want int
	}{
		{in: []int{7, 1, 5, 3, 6, 4}, want: 7},
		{in: []int{1, 2, 3, 4, 5}, want: 4},
		{in: []int{7, 6, 4, 3, 1}, want: 0},
		{in: []int{2, 1, 2, 0, 1}, want: 2},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := maxProfit(tt.in)
			if got != tt.want {
				t.Fatalf("got: %v, want: %v", got, tt.want)
			}
		})
	}

}
