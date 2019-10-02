package removeelement

import (
	"fmt"
	"testing"
)

type Element struct {
	nums []int
	val  int
}

var cases = []struct {
	id    int
	input Element
	want  int
}{
	{id: 1, input: Element{nums: []int{3, 2, 2, 3}, val: 3}, want: 2},
	{id: 2, input: Element{nums: []int{0, 1, 2, 2, 3, 0, 4, 2}, val: 2}, want: 5},
	{id: 3, input: Element{nums: []int{}, val: 2}, want: 0},
	{id: 4, input: Element{nums: []int{1}, val: 2}, want: 1},
	{id: 5, input: Element{nums: []int{2}, val: 2}, want: 0},
}

func TestRemoveElement(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {

			got := removeElement(tt.input.nums, tt.input.val)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
