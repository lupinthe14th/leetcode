package main

// TreeNode is definition for a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	acc := make([]int, 0)
	preorderRecursive(&acc, root)
	return acc
}

func preorderRecursive(acc *[]int, root *TreeNode) {
	if root == nil {
		return
	}
	*acc = append(*acc, root.Val)
	preorderRecursive(acc, root.Left)
	preorderRecursive(acc, root.Right)
}
