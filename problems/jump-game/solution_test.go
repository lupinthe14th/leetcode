package main

import (
	"fmt"
	"testing"
)

var cases = []struct {
	id    int
	input []int
	want  bool
}{
	{id: 1, input: []int{2, 3, 1, 1, 4}, want: true},
	{id: 2, input: []int{3, 2, 1, 0, 4}, want: false},
	{id: 3, input: []int{0}, want: true},
	{id: 4, input: []int{2, 0, 0}, want: true},
	{id: 5, input: []int{2, 5, 0, 0}, want: true},
}

func TestSolution(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := canJump(tt.input)
			if got != tt.want {
				t.Errorf("%t, want: %t", got, tt.want)
			}
		})
	}
}
