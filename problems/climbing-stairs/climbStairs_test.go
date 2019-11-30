package climbingstairs

import (
	"fmt"
	"testing"
)

var cases = []struct {
	id    int
	input int
	want  int
}{
	{id: 1, input: 1, want: 1},
	{id: 2, input: 2, want: 2},
	{id: 3, input: 3, want: 3},
	{id: 4, input: 4, want: 5},
	{id: 5, input: 5, want: 8},
}

func TestClimbStairs(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := climbStairs(tt.input)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
