package linkedlistcycle

import (
	"fmt"
	"testing"
)

type input struct {
	is  []int
	pos int
}

var cases = []struct {
	id    int
	input input
	want  bool
}{
	{id: 1, input: input{is: []int{3, 2, 0, -4}, pos: 1}, want: true},
	{id: 2, input: input{is: []int{1, 2}, pos: 0}, want: true},
	{id: 3, input: input{is: []int{1}, pos: -1}, want: false},
}

func intSlice2CycleListNode(is []int, pos int) *ListNode {
	l, p := &ListNode{}, &ListNode{}
	out := l
	for i, v := range is {
		l.Val = v
		if i == pos {
			p = l
		}
		if i < len(is)-1 {
			l.Next = &ListNode{}
			l = l.Next
		} else {
			l.Next = p
		}
	}
	return out
}

func TestHasCycle(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := hasCycle(intSlice2CycleListNode(tt.input.is, tt.input.pos))
			if got != tt.want {
				t.Errorf("%t, want: %t", got, tt.want)
			}
		})
	}
}
