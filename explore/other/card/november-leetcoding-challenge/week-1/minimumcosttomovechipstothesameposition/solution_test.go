package minimumcosttomovechipstothesameposition

import (
	"fmt"
	"testing"
)

func TestMinCostToMoveChips(t *testing.T) {
	t.Parallel()
	type in struct {
		position, m []int
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{position: []int{1, 2, 3}}, want: 1},
		{in: in{position: []int{2, 2, 2, 3, 3}}, want: 2},
		{in: in{position: []int{1, 1000000000}}, want: 1},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := minCostToMoveChips(tt.in.position)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
