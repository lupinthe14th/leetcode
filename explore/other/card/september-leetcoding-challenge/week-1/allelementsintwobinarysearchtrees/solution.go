package allelementsintwobinarysearchtrees

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	out := make([]int, 0)
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		out = append(out, node.Val)
		if node.Left != nil {
			traversal(node.Left)
		}
		if node.Right != nil {
			traversal(node.Right)
		}
	}
	traversal(root1)
	traversal(root2)
	sort.Ints(out)
	return out
}
