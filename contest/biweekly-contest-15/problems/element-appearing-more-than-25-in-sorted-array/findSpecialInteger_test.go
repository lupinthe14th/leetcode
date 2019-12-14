package elementappearingmorethan25insortedarray

import (
	"fmt"
	"testing"
)

var cases = []struct {
	id    int
	input []int
	want  int
}{
	{id: 1, input: []int{1, 2, 2, 6, 6, 6, 6, 7, 10}, want: 6},
	{id: 2, input: []int{1}, want: 1},
}

func TestFindSpecialInteger(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := findSpecialInteger(tt.input)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
