package rotatelist

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	l := 0
	tail := head
	for tail.Next != nil {
		l++
		tail = tail.Next
	}
	l++
	tail.Next = head
	r := l - k%l
	tail = head
	for i := r; i > 1; i-- {
		tail = tail.Next
	}
	head = tail.Next
	tail.Next = nil

	return head
}
