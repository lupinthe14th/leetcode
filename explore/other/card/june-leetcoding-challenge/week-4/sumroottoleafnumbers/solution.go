package sumroottoleafnumbers

import (
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	a := []string{}
	var number func(root *TreeNode, s string)

	number = func(root *TreeNode, s string) {
		if root == nil {
			return
		}
		v := strconv.Itoa(root.Val)
		if root.Left != nil {
			number(root.Left, s+v)
		}
		if root.Right != nil {
			number(root.Right, s+v)
		}
		if root.Left == nil && root.Right == nil {
			a = append(a, s+v)
		}
		return
	}
	number(root, "")
	out := 0
	for i := range a {
		n, _ := strconv.Atoi(a[i])
		out += n
	}
	return out
}
