package sortlist

import "sort"

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	memo := make([]int, 0)
	tail := head
	for tail != nil {
		memo = append(memo, tail.Val)
		tail = tail.Next
	}
	sort.Ints(memo)
	out := head
	i := 0
	for head != nil {
		head.Val = memo[i]
		head = head.Next
		i++
	}
	return out
}
