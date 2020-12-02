package maximumdepthofbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	out := 0
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	var traversal func(node *TreeNode, c int)
	traversal = func(node *TreeNode, c int) {
		if node == nil {
			out = max(out, c)
			return
		}
		c++
		traversal(node.Left, c)
		traversal(node.Right, c)
	}

	traversal(root, 0)
	return out
}
