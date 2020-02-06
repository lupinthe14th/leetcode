package reducearraysizetothehalf

import (
	"fmt"
	"testing"
)

type Case struct {
	input []int
	want  int
}

var cases = []Case{
	{input: []int{3, 3, 3, 3, 5, 5, 5, 2, 2, 7}, want: 2},
	{input: []int{7, 7, 7, 7, 7, 7}, want: 1},
	{input: []int{1, 9}, want: 1},
	{input: []int{1000, 1000, 3, 7}, want: 1},
	{input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, want: 5},
}

func TestMinSetSize(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := minSetSize(tt.input)
			if got != tt.want {
				t.Errorf("input: %v, got: %d, want: %d", tt.input, got, tt.want)
			}
		})
	}
}
