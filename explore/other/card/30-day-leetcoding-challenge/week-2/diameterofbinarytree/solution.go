package diameterofbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var c int

func diameterOfBinaryTree(root *TreeNode) int {
	c = 1
	depth(root)
	return c - 1
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := depth(root.Left)
	r := depth(root.Right)
	c = max(c, l+r+1)
	return 1 + max(l, r)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
