package possiblebipartition

func possibleBipartition(N int, dislikes [][]int) bool {
	graph := make([][]int, N+1)
	for i := range graph {
		graph[i] = make([]int, 0, 2)

	}

	for _, edge := range dislikes {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}
	color := make(map[int]int, N+1)
	var dfs func(node, c int) bool

	dfs = func(node, c int) bool {
		if col, ok := color[node]; ok {
			return col == c
		}
		color[node] = c

		for _, i := range graph[node] {
			if !dfs(i, c^1) {
				return false
			}
		}
		return true
	}

	for node := 1; node <= N; node++ {
		if _, ok := color[node]; !ok && !dfs(node, 0) {
			return false
		}
	}
	return true
}
