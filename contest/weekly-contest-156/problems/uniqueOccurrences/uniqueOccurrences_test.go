package uniqueoccurrences

import (
	"fmt"
	"testing"
)

var cases = []struct {
	id    int
	input []int
	want  bool
}{
	{id: 1, input: []int{1, 2, 2, 1, 1, 3}, want: true},
	{id: 2, input: []int{1, 2}, want: false},
	{id: 3, input: []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}, want: true},
}

func TestUniqueOccurrences(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprintln(tt.id), func(t *testing.T) {
			got := uniqueOccurrences(tt.input)
			if got != tt.want {
				t.Errorf("%t, want: %t", got, tt.want)
			}
		})
	}
}
