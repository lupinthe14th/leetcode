package besttimetobuyandsellstock

import (
	"fmt"
	"testing"
)

var cases = []struct {
	id    int
	input []int
	want  int
}{
	{id: 1, input: []int{7, 1, 5, 3, 6, 4}, want: 5},
	{id: 2, input: []int{7, 6, 4, 3, 1}, want: 0},
	{id: 3, input: []int{2, 1, 4}, want: 3},
}

func TestMaxProfit(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := maxProfit(tt.input)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
