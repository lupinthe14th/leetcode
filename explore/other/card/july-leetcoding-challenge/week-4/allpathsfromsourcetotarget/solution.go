package allpathsfromsourcetotarget

func allPathsSourceTarget(graph [][]int) [][]int {

	out := make([][]int, 0)

	var dfs func(node int, path []int)
	dfs = func(node int, path []int) {
		if len(graph[node]) == 0 {
			out = append(out, path)
			return
		}
		for _, edge := range graph[node] {
			p := make([]int, len(path))
			copy(p, path)
			p = append(p, edge)
			dfs(edge, p)
		}
	}

	dfs(0, []int{0})
	return out
}
