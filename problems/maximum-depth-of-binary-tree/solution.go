// Refarence: https://leetcode.com/problems/maximum-depth-of-binary-tree/discuss/34216/Simple-solution-using-Java
package maximumdepthofbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func max(x, y int) int {
	if x > y {
		return x

	}
	return y
}
