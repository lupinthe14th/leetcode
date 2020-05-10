package minimumtimetocollectallapplesinatree

// SeeAlso: https://leetcode.com/problems/minimum-time-to-collect-all-apples-in-a-tree/discuss/623673/C++-Java-Simple-code-with-comments-Key-Idea-Algorithm-No-need-of-visited
var ug = make(map[int][]int) // to store the graph

func minTime(n int, edges [][]int, hasApple []bool) int {

	cg(edges)                  // construct the graph first.
	return dfs(0, 0, hasApple) // cost of reaching the root is 0. For all others, its 2.
}

func cg(edges [][]int) {
	ug = map[int][]int{} // init graph
	for _, edge := range edges {
		ug[edge[0]] = append(ug[edge[0]], edge[1]) // adjecency list representation
	}
}

func dfs(node, cost int, hasApple []bool) int {
	childrenCost := 0 // cost of traversing all children.
	for _, x := range ug[node] {
		childrenCost += dfs(x, 2, hasApple) // check recursively for all apples.
	}
	if childrenCost == 0 && hasApple[node] == false {
		// If no child has apples, then we won't traverse the subtree, so cost will be zero.
		// similarly, if current node also does not have the apple, we won't traverse this branch at all, so cost will be zero.
		return 0
	}
	// Children has at least one apple or the current node has an apple, so add those costs.
	return childrenCost + cost
}
