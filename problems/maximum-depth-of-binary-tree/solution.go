package maximumdepthofbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ans int

func maxDepth(root *TreeNode) int {
	ans = 0
	maximumDepth(root, 1)
	return ans
}

func maximumDepth(root *TreeNode, depth int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		ans = max(ans, depth)
	}
	maximumDepth(root.Left, depth+1)
	maximumDepth(root.Right, depth+1)
}

func max(x, y int) int {
	if x > y {
		return x

	}
	return y
}
