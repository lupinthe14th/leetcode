package linkedlistinbinarytree

import (
	"fmt"
	"testing"
)

type in struct {
	head *ListNode
	root *TreeNode
}

type Case struct {
	in   in
	want bool
}

var cases = []Case{
	{in: in{
		head: &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 2, Next: &ListNode{Val: 6}}}},
		root: &TreeNode{Val: 1, Left: &TreeNode{Val: 4, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 6}}, Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}}}}, want: true},
}

func TestIsSubPath(t *testing.T) {
	for i, c := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := isSubPath(c.in.head, c.in.root)
			if got != c.want {
				t.Errorf("in: %#+v, got: %t, want: %t", c.in, got, c.want)
			}
		})
	}
}
