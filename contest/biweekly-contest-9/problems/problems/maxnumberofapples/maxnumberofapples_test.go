package maxnumberofapples

import (
	"fmt"
	"testing"
)

var cases = []struct {
	id    int
	input []int
	want  int
}{
	{id: 1, input: []int{100, 200, 150, 1000}, want: 4},
	{id: 2, input: []int{900, 950, 800, 1000, 700, 800}, want: 5},
}

func TestMaxNumberOfApples(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := maxNumberOfApples(tt.input)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
