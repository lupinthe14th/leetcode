package invertbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	r := invertTree(root.Right)
	l := invertTree(root.Left)
	root.Left = r
	root.Right = l
	return root
}
