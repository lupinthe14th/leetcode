package maximumdepthofbinarytree

import (
	"fmt"
	"testing"
)

type Case struct {
	in   *TreeNode
	want int
}

var cases = []Case{
	{in: &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}, want: 3},
	{in: &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15, Right: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 7}}}, want: 4},
	{in: nil, want: 0},
}

func TestMaxDepth(t *testing.T) {
	for i, c := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := maxDepth(c.in)
			if got != c.want {
				t.Errorf("%d, want: %d", got, c.want)
			}
		})
	}
}
