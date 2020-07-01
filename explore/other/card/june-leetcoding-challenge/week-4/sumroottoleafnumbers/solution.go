package sumroottoleafnumbers

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	var sumNode func(root *TreeNode, p int) int

	sumNode = func(root *TreeNode, p int) int {
		sum := 0
		if root == nil {
			return 0
		}
		if root.Left != nil {
			sum = sum + sumNode(root.Left, p*10+root.Val)
		}
		if root.Right != nil {
			sum = sum + sumNode(root.Right, p*10+root.Val)
		}
		if root.Left == nil && root.Right == nil {
			sum = p*10 + root.Val
		}
		return sum
	}
	return sumNode(root, 0)
}
