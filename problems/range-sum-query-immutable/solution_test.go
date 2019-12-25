package rangesumqueryimmutable

import (
	"fmt"
	"testing"
)

type input struct {
	nums []int
	i, j int
}

var cases = []struct {
	id    int
	input input
	want  int
}{
	{id: 1, input: input{nums: []int{-2, 0, 3, -5, 2, -1}, i: 0, j: 2}, want: 1},
	{id: 2, input: input{nums: []int{-2, 0, 3, -5, 2, -1}, i: 2, j: 5}, want: -1},
	{id: 3, input: input{nums: []int{-2, 0, 3, -5, 2, -1}, i: 0, j: 5}, want: -3},
}

func TestSumRange(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			c := Constructor(tt.input.nums)
			got := c.SumRange(tt.input.i, tt.input.j)
			if got != tt.want {
				t.Errorf("%d, want: %d (%+v)", got, tt.want, tt.input)
			}
		})
	}
}
