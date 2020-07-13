package sametree

import "reflect"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return reflect.DeepEqual(p, q)
}
