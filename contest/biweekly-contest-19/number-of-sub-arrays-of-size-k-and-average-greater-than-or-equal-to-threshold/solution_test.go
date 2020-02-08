package numberofsubarraysofsizekandaveragegreaterthanorequaltothreshold

import (
	"fmt"
	"testing"
)

type Case struct {
	input input
	want  int
}

type input struct {
	arr       []int
	k         int
	threshold int
}

var cases = []Case{
	{input: input{arr: []int{2, 2, 2, 2, 5, 5, 5, 8}, k: 3, threshold: 4}, want: 3},
	{input: input{arr: []int{1, 1, 1, 1, 1}, k: 1, threshold: 0}, want: 5},
	{input: input{arr: []int{11, 13, 17, 23, 29, 31, 7, 5, 2, 3}, k: 3, threshold: 5}, want: 6},
	{input: input{arr: []int{7, 7, 7, 7, 7, 7, 7}, k: 7, threshold: 7}, want: 1},
	{input: input{arr: []int{4, 4, 4, 4}, k: 4, threshold: 1}, want: 1},
}

func TestNumOfSubarrays(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := numOfSubarrays(tt.input.arr, tt.input.k, tt.input.threshold)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
