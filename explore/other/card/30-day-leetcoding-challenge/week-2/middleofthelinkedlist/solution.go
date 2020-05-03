package middleofthelinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

// refactor
// approach 2: Fast and Slow Pointer
// Intuition and Algorithm
//
// When traversing the list with a pointer slow, make another pointer fast that
// traverses twice as fast. When fast reaches the end of the list, slow must be
// in the middle.
func middleNode(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
