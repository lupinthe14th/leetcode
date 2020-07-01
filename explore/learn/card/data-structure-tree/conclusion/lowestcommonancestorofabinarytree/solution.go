package lowestcommonancestorofabinarytree

import "reflect"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	out := &TreeNode{}
	var recurseTree func(cur *TreeNode) bool
	recurseTree = func(cur *TreeNode) bool {
		if cur == nil {
			return false
		}
		left := 0
		if recurseTree(cur.Left) {
			left = 1
		}
		right := 0
		if recurseTree(cur.Right) {
			right = 1
		}
		mid := 0
		if reflect.DeepEqual(cur, p) || reflect.DeepEqual(cur, q) {
			mid = 1
		}

		if mid+left+right >= 2 {
			out = cur
		}
		return mid+left+right > 0
	}

	recurseTree(root)
	return out
}
