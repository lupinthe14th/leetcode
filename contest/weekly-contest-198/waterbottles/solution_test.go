package waterbottles

import (
	"fmt"
	"testing"
)

func TestNumWaterBottles(t *testing.T) {
	t.Parallel()
	type in struct {
		numBottles, numExchange int
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{numBottles: 9, numExchange: 3}, want: 13},
		{in: in{numBottles: 15, numExchange: 4}, want: 19},
		{in: in{numBottles: 5, numExchange: 5}, want: 6},
		{in: in{numBottles: 2, numExchange: 3}, want: 2},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := numWaterBottles(tt.in.numBottles, tt.in.numExchange)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
