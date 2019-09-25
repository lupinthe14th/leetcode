package mergetwosortedlists

import (
	"sort"
)

// ListNode is singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var stack []int
	var ln *ListNode
	for l := l1; l != nil; l = l.Next {
		stack = append(stack, l.Val)
	}

	for l := l2; l != nil; l = l.Next {
		stack = append(stack, l.Val)
	}

	sort.Ints(stack)

	for i := len(stack); i > 0; i-- {
		if i == len(stack) {
			ln = &ListNode{Val: stack[i-1], Next: nil}
		} else {
			ln = &ListNode{Val: stack[i-1], Next: ln}
		}
	}
	return ln
}
