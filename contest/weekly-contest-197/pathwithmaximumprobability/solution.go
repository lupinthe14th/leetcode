package pathwithmaximumprobability

func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {

	type state struct {
		node int
		prob float64
	}

	memo := make(map[int][][]float64)
	for i, edge := range edges {
		if _, ok := memo[edge[0]]; !ok {
			memo[edge[0]] = make([][]float64, 0)
		}
		if _, ok := memo[edge[1]]; !ok {
			memo[edge[1]] = make([][]float64, 0)
		}
		memo[edge[0]] = append(memo[edge[0]], []float64{float64(edge[1]), succProb[i]})
		memo[edge[1]] = append(memo[edge[1]], []float64{float64(edge[0]), succProb[i]})

	}
	q := make([]state, n)
	probs := make([]float64, n)
	q[0] = state{node: start, prob: 1.0}
	for len(q) != 0 {
		s := q[0]
		q = q[1:]
		parent := s.node
		prob := s.prob

		childs, ok := memo[parent]
		if !ok {
			childs = make([][]float64, 0)
		}
		for _, child := range childs {
			if probs[int(child[0])] >= prob*child[1] {
				continue
			}
			q = append(q, state{int(child[0]), prob * child[1]})
			probs[int(child[0])] = prob * child[1]
		}
	}
	return probs[end]
}
