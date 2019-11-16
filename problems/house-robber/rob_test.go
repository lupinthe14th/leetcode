package houserobber

import (
	"fmt"
	"testing"
)

var cases = []struct {
	id    int
	input []int
	want  int
}{
	{id: 1, input: []int{1, 2, 3, 1}, want: 4},
	{id: 2, input: []int{2, 7, 9, 3, 1}, want: 12},
	{id: 3, input: []int{}, want: 0},
	{id: 4, input: []int{0}, want: 0},
	{id: 5, input: []int{4}, want: 4},
	{id: 6, input: []int{2, 1, 1, 2}, want: 4},
}

func TestRob(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := rob(tt.input)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
