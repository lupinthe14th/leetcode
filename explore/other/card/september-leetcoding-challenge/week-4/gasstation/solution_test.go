package gasstation

import (
	"fmt"
	"testing"
)

func TestCanCompleteCircuit(t *testing.T) {
	t.Parallel()
	type in struct {
		gas, cost []int
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{gas: []int{1, 2, 3, 4, 5}, cost: []int{3, 4, 5, 1, 2}}, want: 3},
		{in: in{gas: []int{2, 3, 4}, cost: []int{3, 4, 3}}, want: -1},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := canCompleteCircuit(tt.in.gas, tt.in.cost)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
