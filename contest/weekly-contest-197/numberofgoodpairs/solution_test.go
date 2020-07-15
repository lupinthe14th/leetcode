package numberofgoodpairs

import (
	"fmt"
	"testing"
)

func TestNumIdenticalPairs(t *testing.T) {
	tests := []struct {
		in   []int
		want int
	}{
		{in: []int{1, 2, 3, 1, 1, 3}, want: 4},
		{in: []int{1, 1, 1, 1}, want: 6},
		{in: []int{1, 2, 3}, want: 0},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := numIdenticalPairs(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
