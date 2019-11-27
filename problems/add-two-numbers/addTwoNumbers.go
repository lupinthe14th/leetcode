package main

import (
	"math"
)

// ListNode is singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseOrder(l *ListNode) int {
	var ans int
	i := 0
	for l != nil {
		ans += l.Val * int(math.Pow(10, float64(i)))
		l = l.Next
		i++
	}
	return ans
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ln := &ListNode{}
	out := ln
	sum := reverseOrder(l1) + reverseOrder(l2)

	for sum != 0 {
		ln.Val = sum % 10
		sum /= 10
		if sum != 0 {
			ln.Next = &ListNode{}
			ln = ln.Next
		}
	}
	return out
}
