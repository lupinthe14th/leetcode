package increasingordersearchtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {
	var memo []int
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		memo = append(memo, node.Val)
		traversal(node.Right)
	}

	traversal(root)

	out := &TreeNode{Val: memo[0]}
	tmp := out
	for i := 1; i < len(memo); i++ {
		tmp.Right = &TreeNode{Val: memo[i]}
		tmp = tmp.Right
	}
	return out
}
