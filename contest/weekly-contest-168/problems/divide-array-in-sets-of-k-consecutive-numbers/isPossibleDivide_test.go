package dividearrayinsetsofkconsecutivenumbers

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
	{id: 1, input: input{nums: []int{1, 2, 3, 3, 4, 4, 5, 6}, k: 4}, want: true},
	{id: 2, input: input{nums: []int{3, 2, 1, 2, 3, 4, 3, 4, 5, 9, 10, 11}, k: 3}, want: true},
	{id: 3, input: input{nums: []int{3, 3, 2, 2, 1, 1}, k: 3}, want: true},
	{id: 4, input: input{nums: []int{1, 2, 3, 4}, k: 3}, want: false},
	{id: 5, input: input{nums: []int{15, 16, 17, 18, 19, 16, 17, 18, 19, 20, 6, 7, 8, 9, 10, 3, 4, 5, 6, 20}, k: 5}, want: false},
	{id: 6, input: input{nums: []int{15, 16, 17, 18, 19, 16, 17, 18, 19, 20, 6, 7, 8, 9, 10, 3, 4, 5, 6, 20}, k: 1}, want: true},
}

func TestIsPossibleDivide(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := isPossibleDivide(tt.input.nums, tt.input.k)
			if got != tt.want {
				t.Errorf("%t, want: %t", got, tt.want)
			}
		})
	}
}
