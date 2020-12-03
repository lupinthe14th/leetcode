package linkedlistrandomnode

import (
	"math/rand"
	"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	r []int
}

/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
	r := make([]int, 0)
	for head != nil {
		r = append(r, head.Val)
		head = head.Next
	}
	return Solution{r: r}
}

/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
	rand.Seed(time.Now().UnixNano())
	pick := rand.Intn(len(this.r))
	return this.r[pick]
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
