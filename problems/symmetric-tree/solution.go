package main

// TreeNode is Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
}

func isMirror(r, l *TreeNode) bool {
	if r == nil && l == nil {
		return true
	}
	if r == nil || l == nil {
		return false
	}
	return (r.Val == l.Val) && isMirror(r.Right, l.Left) && isMirror(r.Left, l.Right)
}

func main() {}
