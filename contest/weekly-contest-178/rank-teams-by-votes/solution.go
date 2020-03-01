// reference:
// - https://leetcode.com/problems/rank-teams-by-votes/discuss/524869/Python-Count-Votes
// - https://leetcode.com/problems/rank-teams-by-votes/discuss/525210/Go-sort-multiple-votes
package rankteamsbyvotes

import (
	"sort"
)

func rankTeams(votes []string) string {
	l := len(votes[0])
	count := make(map[rune][]int)

	for _, v := range votes[0] {
		count[v] = make([]int, l)
	}

	for _, vote := range votes {
		for i, v := range vote {
			count[v][i] += 1
		}
	}
	r := []rune(votes[0])
	sort.Slice(r, func(i, j int) bool {
		for k := 0; k < l; k++ {
			if count[r[i]][k] != count[r[j]][k] {
				return count[r[i]][k] > count[r[j]][k]
			}
		}
		return r[i] < r[j]
	})
	return string(r)
}
