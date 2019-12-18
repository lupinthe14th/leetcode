package subarraysumequalsk

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
	want  int
}{
	{id: 1, input: input{nums: []int{1, 1, 1}, k: 2}, want: 2},
}

func TestSubarraySum(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := subarraySum(tt.input.nums, tt.input.k)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
