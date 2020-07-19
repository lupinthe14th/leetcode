package numberofnodesinthesubtreewiththesamelabel

func countSubTrees(n int, edges [][]int, labels string) []int {
	out := make([]int, n)

	adjList := make(map[int][]int)

	// create adjlist
	for i := range edges {
		src := edges[i][0]
		dist := edges[i][1]
		adjList[src] = append(adjList[src], dist)
		adjList[dist] = append(adjList[dist], src)
	}

	visit := make([]bool, n)
	var dfs func(cur int) [26]int
	dfs = func(cur int) [26]int {
		visit[cur] = true
		var v [26]int
		v[labels[cur]-'a'] = 1
		if neighbors, ok := adjList[cur]; ok {
			for _, neighbor := range neighbors {
				if visit[neighbor] {
					continue
				}

				r := dfs(neighbor)
				for i := range r {
					v[i] += r[i]
				}
			}
		}
		out[cur] = v[labels[cur]-'a']
		return v
	}
	dfs(0)
	return out
}
