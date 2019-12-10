package mincostclimbingstairs

import (
	"fmt"
	"testing"
)

var cases = []struct {
	id    int
	input []int
	want  int
}{
	{id: 1, input: []int{10, 15, 20}, want: 15},
	{id: 2, input: []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, want: 6},
}

func TestMinCostClimbingStairs(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := minCostClimbingStairs(tt.input)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
