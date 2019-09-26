package removeduplicatesfromsortedarray

import (
	"fmt"
	"testing"
)

var cases = []struct {
	id    int
	input []int
	want  int
}{
	{id: 1, input: []int{1, 1, 2}, want: 2},
	{id: 2, input: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, want: 5},
}

func TestRemoveDupulicates(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := removeDuplicates(tt.input)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
