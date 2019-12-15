package convertbinarynumberinalinkedlisttointeger

import (
	"strconv"
)

// ListNode is Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {

	var s string

	for head != nil {
		s += strconv.Itoa(head.Val)
		head = head.Next
	}
	i, _ := strconv.ParseInt(s, 2, 0)
	return int(i)
}
