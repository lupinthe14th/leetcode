package findwinneronatictactoegame

import (
	"fmt"
	"testing"
)

var cases = []struct {
	id int

	input [][]int
	want  string
}{
	{id: 1, input: [][]int{{0, 0}, {2, 0}, {1, 1}, {2, 1}, {2, 2}}, want: "A"},
	{id: 2, input: [][]int{{0, 0}, {1, 1}, {0, 1}, {0, 2}, {1, 0}, {2, 0}}, want: "B"},
	{id: 3, input: [][]int{{0, 0}, {1, 1}, {2, 0}, {1, 0}, {1, 2}, {2, 1}, {0, 1}, {0, 2}, {2, 2}}, want: "Draw"},
	{id: 4, input: [][]int{{0, 0}, {1, 1}}, want: "Pending"},
	{id: 5, input: [][]int{{1, 0}, {2, 2}, {2, 0}, {0, 1}, {1, 1}}, want: "Pending"},
}

func TestTictactoe(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := tictactoe(tt.input)
			if got != tt.want {
				t.Errorf("%s, want: %s", got, tt.want)
			}
		})
	}
}
