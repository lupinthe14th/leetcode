package sumroottoleafnumbers

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	a := []int{}
	var number func(root *TreeNode, p int)

	number = func(root *TreeNode, p int) {
		if root == nil {
			return
		}
		if root.Left != nil {
			number(root.Left, p*10+root.Val)
		}
		if root.Right != nil {
			number(root.Right, p*10+root.Val)
		}
		if root.Left == nil && root.Right == nil {
			a = append(a, p*10+root.Val)
		}
		return
	}
	number(root, 0)
	out := 0
	for i := range a {
		out += a[i]
	}
	return out
}
