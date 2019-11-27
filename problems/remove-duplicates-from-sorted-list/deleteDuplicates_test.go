package removeduplicatesfromsortedlist

import (
	"fmt"
	"reflect"
	"testing"
)

var cases = []struct {
	id    int
	input []int
	want  []int
}{
	{id: 1, input: []int{1, 1, 2}, want: []int{1, 2}},
	{id: 2, input: []int{1, 1, 2, 3, 3}, want: []int{1, 2, 3}},
	{id: 3, input: []int{0, 0, 0, 0, 0}, want: []int{0}},
}

func intSlice2ListNode(is []int) *ListNode {
	l := &ListNode{}
	out := l

	for i, v := range is {
		l.Val = v

		if i < len(is)-1 {
			l.Next = &ListNode{}
			l = l.Next
		}
	}
	return out
}

func listNode2IntSlice(l *ListNode) []int {
	var is []int

	for l != nil {
		is = append(is, l.Val)
		l = l.Next
	}
	return is
}

func TestDeleteDuplicates(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := listNode2IntSlice(deleteDuplicates(intSlice2ListNode(tt.input)))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%v, want: %v", got, tt.want)
			}
		})
	}
}
