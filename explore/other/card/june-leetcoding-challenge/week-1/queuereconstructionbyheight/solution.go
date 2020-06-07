package queuereconstructionbyheight

import (
	"sort"
)

// https://leetcode.com/problems/queue-reconstruction-by-height/discuss/577514/Go-12ms-97-solution
func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] > people[j][1]
		}
		return people[i][0] < people[j][0]
	})
	idxs := make([]int, len(people))
	for i := range idxs {
		idxs[i] = i
	}
	out := make([][]int, len(people))
	for _, v := range people {
		out[idxs[v[1]]] = v
		idxs = append(idxs[:v[1]], idxs[v[1]+1:]...)
	}
	return out
}
