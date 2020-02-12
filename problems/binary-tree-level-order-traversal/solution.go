// refarence: https://leetcode.com/problems/binary-tree-level-order-traversal/discuss/33463/Python-easy-to-understand-iterative-queue-solution
package binarytreelevelordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Queue struct {
	node *TreeNode
	idx  int
}

func levelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}

	var idx int
	var queue = []Queue{
		{node: root, idx: idx},
	}

	for len(queue) != 0 {
		visit := queue[0].node
		level := queue[0].idx
		queue = queue[1:]

		if len(ans) == level {
			ans = append(ans, []int{visit.Val})
		} else {
			ans[level] = append(ans[level], visit.Val)
		}
		if visit.Left != nil {
			queue = append(queue, Queue{node: visit.Left, idx: level + 1})
		}

		if visit.Right != nil {
			queue = append(queue, Queue{node: visit.Right, idx: level + 1})
		}

	}
	return ans
}
