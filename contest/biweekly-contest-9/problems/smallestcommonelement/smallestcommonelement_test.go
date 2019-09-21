package smallestcommonelement

import (
	"fmt"
	"testing"
)

var cases = []struct {
	id    int
	input [][]int
	want  int
}{
	{id: 1, input: [][]int{{1, 2, 3, 4, 5}, {2, 4, 5, 8, 10}, {3, 5, 7, 9, 11}, {1, 3, 5, 7, 9}}, want: 5},
}

func TestSmallestCommonElement(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := smallestCommonElement(tt.input)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
