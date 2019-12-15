package convertbinarynumberinalinkedlisttointeger

import (
	"fmt"
	"testing"
)

var cases = []struct {
	id    int
	input []int
	want  int
}{
	{id: 1, input: []int{1, 0, 1}, want: 5},
	{id: 2, input: []int{0}, want: 0},
	{id: 3, input: []int{1}, want: 1},
	{id: 4, input: []int{1, 0, 0, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0}, want: 18880},
	{id: 5, input: []int{0, 0}, want: 0},
}

func intSlice2ListNode(is []int) *ListNode {
	l := &ListNode{}
	out := l

	if len(is) == 0 {
		return nil
	}
	for i, v := range is {
		l.Val = v

		if i < len(is)-1 {
			l.Next = &ListNode{}
			l = l.Next
		}
	}
	return out
}

func TestGetDecimalValue(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := getDecimalValue(intSlice2ListNode(tt.input))
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
