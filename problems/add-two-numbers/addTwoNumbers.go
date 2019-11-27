package main

import (
	"math"
)

// ListNode is singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseOrder(l *ListNode) int64 {
	var ans int64
	i := 0
	for l != nil {
		ans += int64(l.Val * int(math.Pow(10, float64(i))))
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
		ln.Val = int(sum % 10)
		sum /= 10
		if sum != 0 {
			ln.Next = &ListNode{}
			ln = ln.Next
		}
	}
	return out
}
