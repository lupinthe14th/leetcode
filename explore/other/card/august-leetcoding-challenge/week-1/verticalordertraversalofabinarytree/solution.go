package verticalordertraversalofabinarytree

import (
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// SeeAlso:
// - https://leetcode.com/problems/vertical-order-traversal-of-a-binary-tree/discuss/659293/Super-simple-easy-to-understand-fast-basic-recursive-traversal-solution
// - https://leetcode.com/problems/vertical-order-traversal-of-a-binary-tree/discuss/231113/C++-hashmap-vs.-map
// - https://leetcode.com/problems/vertical-order-traversal-of-a-binary-tree/discuss/449743/Go-Level-order-traversal-with-memorization
func verticalTraversal(root *TreeNode) [][]int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	levels := make(map[int]map[int][]int)
	var deep int
	var traversal func(node *TreeNode, w, d int)
	traversal = func(node *TreeNode, w, d int) {
		if node == nil {
			return
		}

		deep = max(deep, d)
		levelNodes, ok := levels[w]
		if !ok {
			levelNodes = make(map[int][]int)
		}
		levelNodes[d] = append(levelNodes[d], node.Val)
		levels[w] = levelNodes

		if node.Left != nil {
			traversal(node.Left, w-1, d+1)
		}

		if node.Right != nil {
			traversal(node.Right, w+1, d+1)
		}
	}

	traversal(root, 0, 0)
	out := make([][]int, 0)
	for i := -1000; i <= 1000; i++ {
		level, ok := levels[i]
		if !ok {
			continue
		}
		levelVals := make([]int, 0)
		for j := 0; j <= deep; j++ {
			vals, ok := level[j]
			if !ok {
				continue
			}
			sort.Ints(vals)
			levelVals = append(levelVals, vals...)
		}
		out = append(out, levelVals)
	}
	return out
}
