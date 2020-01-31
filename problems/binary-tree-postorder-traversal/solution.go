package binarytreepostordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	acc := make([]int, 0)
	postorderTraversalRecursive(root, &acc)
	return acc
}

func postorderTraversalRecursive(root *TreeNode, acc *[]int) {
	if root == nil {
		return
	}
	postorderTraversalRecursive(root.Left, acc)
	postorderTraversalRecursive(root.Right, acc)
	*acc = append(*acc, root.Val)
}
