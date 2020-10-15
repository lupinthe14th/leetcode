package houserobberii

import (
	"fmt"
	"testing"
)

func TestRob(t *testing.T) {
	t.Parallel()
	type in struct {
		nums, m []int
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{nums: []int{2, 3, 2}}, want: 3},
		{in: in{nums: []int{1, 2, 3, 1}}, want: 4},
		{in: in{nums: []int{0}}, want: 0},
		{in: in{nums: []int{1}}, want: 1},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := rob(tt.in.nums)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
