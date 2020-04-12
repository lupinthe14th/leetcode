package queriesonapermutationwithkey

func processQueries(queries []int, m int) []int {
	p := make([]int, m)
	for i := 0; i < m; i++ {
		p[i] = i + 1
	}

	ans := make([]int, len(queries))
	for i := range queries {
		// return idx
		idx := 0
		for idx = range p {
			if queries[i] == p[idx] {
				break
			}
		}
		//delete
		p = append(p[:idx], p[idx+1:m]...)
		//push
		p = append([]int{queries[i]}, p...)
		ans[i] = idx
	}
	return ans
}
