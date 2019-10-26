package maximumsubarray

import (
	"fmt"
	"testing"
)

var cases = []struct {
	id    int
	input []int
	want  int
}{
	{id: 1, input: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, want: 6},
}

func TestMaxSubArray(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {

			got := maxSubArray(tt.input)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
