// https://leetcode.com/problems/binary-tree-maximum-path-sum/discuss/39775/Accepted-short-solution-in-Java
package binarytreemaximumpathsum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ans int

func maxPathSum(root *TreeNode) int {
	ans = -1 << 31

	_ = dfs(root)
	return ans
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := max(0, dfs(root.Left))
	right := max(0, dfs(root.Right))

	ans = max(ans, left+right+root.Val)
	return max(left, right) + root.Val
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
