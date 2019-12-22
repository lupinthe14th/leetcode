package continuoussubarraysum

import (
	"fmt"
	"testing"
)

type input struct {
	nums []int
	k    int
}

var cases = []struct {
	id    int
	input input
	want  bool
}{
	{id: 1, input: input{nums: []int{23, 2, 4, 6, 7}, k: 6}, want: true},
	{id: 2, input: input{nums: []int{23, 2, 6, 4, 7}, k: 6}, want: true},
	{id: 3, input: input{nums: []int{23, 2, 4, 6, 7}, k: 0}, want: false},
	{id: 4, input: input{nums: []int{0, 0}, k: 0}, want: true},
}

func TestCheckSubarraySum(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := checkSubarraySum(tt.input.nums, tt.input.k)
			if got != tt.want {
				t.Errorf("%t, want: %t", got, tt.want)
			}
		})
	}
}
