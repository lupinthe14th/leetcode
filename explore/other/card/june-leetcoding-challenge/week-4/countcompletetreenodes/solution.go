package countcompletetreenodes

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {

	c := 0

	var count func(root *TreeNode)

	count = func(root *TreeNode) {
		if root == nil {
			return
		}
		c++
		if root.Left != nil {
			count(root.Left)
		}

		if root.Right != nil {
			count(root.Right)
		}
		return
	}
	count(root)
	return c
}
