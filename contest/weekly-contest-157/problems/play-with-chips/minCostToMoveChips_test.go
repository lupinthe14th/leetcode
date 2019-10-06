package playwithchips

import (
	"fmt"
	"testing"
)

var cases = []struct {
	id    int
	input []int
	want  int
}{
	{id: 1, input: []int{2}, want: 0}, //1:0
	{id: 2, input: []int{2, 2}, want: 1},
	{id: 3, input: []int{1, 2, 3}, want: 1}, // 3:1
	{id: 4, input: []int{2, 2, 2, 3}, want: 2},
	{id: 5, input: []int{2, 2, 2, 3, 3}, want: 2}, // 5:2
	{id: 6, input: []int{2, 2, 2, 3, 3, 4}, want: 3},
	{id: 7, input: []int{2, 2, 2, 3, 3, 4, 4}, want: 3}, // 7:3
	{id: 8, input: []int{2, 2, 2, 3, 3, 4, 4, 4}, want: 4},
	{id: 9, input: []int{2, 2, 2, 3, 3, 4, 4, 4, 4}, want: 4}, // 9:4
	{id: 10, input: []int{2, 2, 2, 3, 3, 4, 4, 4, 4, 5}, want: 5},
}

func TestMinCostToMoveChips(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := minCostToMoveChips(tt.input)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
