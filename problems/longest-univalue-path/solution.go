package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ans int

// Algorithm
// Let arrowLength(root) be the length of the longest arrow that extends from the root. That will be 1 + arrowLength(root.Left) if root.Left exists and has the same value as root. Similarly for the root.Right case.
// While we are computing arrow lengths, each candidate answer will be the sum of the arrows in both direstions from that root. We record these candidate answers and return the best one.
func longestUnivaluePath(root *TreeNode) int {
	ans = 0
	arrowLength(root)
	return ans
}

func arrowLength(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := arrowLength(root.Left)
	right := arrowLength(root.Right)
	var arrowLeft, arrowRight int
	if root.Left != nil && root.Left.Val == root.Val {
		arrowLeft += left + 1
	}
	if root.Right != nil && root.Right.Val == root.Val {
		arrowRight += right + 1
	}
	ans = max(ans, arrowLeft+arrowRight)
	return max(arrowLeft, arrowRight)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y

}
func main() {}
