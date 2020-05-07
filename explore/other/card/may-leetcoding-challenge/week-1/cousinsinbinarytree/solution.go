package cousinsinbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// SeeAlso: https://leetcode.com/problems/cousins-in-binary-tree/discuss/501788/Go-BFS-0ms
func isCousins(root *TreeNode, x int, y int) bool {
	return bfs([]*TreeNode{root}, x, y)
}

func bfs(nodes []*TreeNode, x, y int) bool {
	if len(nodes) == 0 {
		return false
	}

	xp, yp, next := -1, -1, []*TreeNode{}

	for _, node := range nodes {
		if node.Left != nil {
			if node.Left.Val == x {
				xp = node.Val
			} else if node.Left.Val == y {
				yp = node.Val
			}
			next = append(next, node.Left)
		}
		if node.Right != nil {
			if node.Right.Val == x {
				xp = node.Val
			} else if node.Right.Val == y {
				yp = node.Val
			}
			next = append(next, node.Right)
		}
	}
	if xp != -1 || yp != -1 {
		return xp != yp && xp != -1 && yp != -1
	}
	return bfs(next, x, y)
}
