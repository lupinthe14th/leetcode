package linkedlistcycle

import (
	"fmt"
	"testing"
)

type input struct {
	is  []int
	pos int
}

type Case struct {
	input input
	want  bool
}

var cases = []Case{
	{input: input{is: []int{3, 2, 0, -4}, pos: 1}, want: true},
	{input: input{is: []int{1, 2}, pos: 0}, want: true},
	{input: input{is: []int{1}, pos: -1}, want: false},
	{input: input{is: []int{}, pos: 0}, want: false},
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
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := hasCycle(intSlice2CycleListNode(tt.input.is, tt.input.pos))
			if got != tt.want {
				t.Errorf("%t, want: %t", got, tt.want)
			}
		})
	}
}
