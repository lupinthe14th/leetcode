package binarytreepostordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	sl := make([]int, 0)
	postorderTraversalRecursive(root, &sl)
	return sl
}

func postorderTraversalRecursive(root *TreeNode, sl *[]int) {
	if root == nil {
		return
	}
	postorderTraversalRecursive(root.Left, sl)
	postorderTraversalRecursive(root.Right, sl)
	*sl = append(*sl, root.Val)
}
