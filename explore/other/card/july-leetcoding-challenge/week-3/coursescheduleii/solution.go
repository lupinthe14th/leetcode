package coursescheduleii

func findOrder(numCourses int, prerequisites [][]int) []int {
	// 最終的に正しいトポロジカルソートを保持するキュー
	order := make([]int, numCourses)

	// 隣接リスト
	adjList := make(map[int][]int)
	// 入力辺の数を保持するリスト
	indegree := make([]int, numCourses)

	// Create adjList
	for i := range prerequisites {
		dest := prerequisites[i][0]
		src := prerequisites[i][1]
		adjList[src] = append(adjList[src], dest)

		indegree[dest]++
	}
	// 次に処理するノード番号を保持するキュー
	processNext := make([]int, 0)
	// indegree[i] == 0となるノード番号iをすべてprocessNextキューに追加
	for i := range indegree {
		if indegree[i] == 0 {
			processNext = append(processNext, i)
		}
	}

	// BFS
	i := 0
	for len(processNext) != 0 {
		cur := processNext[0]
		processNext = processNext[1:]
		// 現在処理ノード番号の出力辺リストがあれば、その出力辺のノード番号に対して入力辺の数をデクリメント
		// 0の値があればそのノード番号を次に処理するキューに格納
		if neighbors, ok := adjList[cur]; ok {
			for _, neighbor := range neighbors {
				indegree[neighbor]--
				if indegree[neighbor] == 0 {
					processNext = append(processNext, neighbor)
				}
			}
		}
		// 現在処理ノードをソートスライスに格納
		order[i] = cur
		i++
	}
	// トポロジカルソートが成立しているか確認
	if i == numCourses {
		return order
	}
	return []int{}
}
