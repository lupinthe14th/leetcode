package sumofroottoleafbinarynumbers

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumRootToLeaf(root *TreeNode) int {

	out := 0
	var traversal func(node *TreeNode, b int)
	traversal = func(node *TreeNode, b int) {
		if node == nil {
			return
		}

		b = b<<1 + node.Val
		if node.Left != nil {
			traversal(node.Left, b)
		}
		if node.Right != nil {
			traversal(node.Right, b)
		}
		if node.Left == nil && node.Right == nil {
			out += b
		}
	}
	traversal(root, 0)
	return out
}
