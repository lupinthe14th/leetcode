package sumofroottoleafbinarynumbers

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumRootToLeaf(root *TreeNode) int {
	bs := make([]int, 0)

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
			bs = append(bs, b)
		}
	}
	traversal(root, 0)
	out := 0
	for i := range bs {
		out += bs[i]
	}
	return out
}
