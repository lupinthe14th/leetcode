package main

// TreeNode is definition for a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)

	for len(stack) != 0 {
		visit := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, visit.Val)
		if visit.Right != nil {
			stack = append(stack, visit.Right)
		}
		if visit.Left != nil {
			stack = append(stack, visit.Left)
		}
	}
	return ans
}

func main() {}
