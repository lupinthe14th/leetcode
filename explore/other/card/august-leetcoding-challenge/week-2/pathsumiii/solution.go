package pathsumiii

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	var helper func(node *TreeNode, cur int) int
	helper = func(node *TreeNode, cur int) int {
		if node == nil {
			return 0
		}

		cur += node.Val
		c := 0
		if cur == sum {
			c++
		}
		return c + helper(node.Left, cur) + helper(node.Right, cur)
	}

	return helper(root, 0) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}
