package reconstructitinerary

import (
	"sort"
)

// refarence: https://leetcode.com/problems/reconstruct-itinerary/discuss/492190/Golang-Topological-sort
func findItinerary(tickets [][]string) []string {
	memo := make(map[string][]string, len(tickets))

	for i := range tickets {
		memo[tickets[i][0]] = append(memo[tickets[i][0]], tickets[i][1])
	}
	for _, v := range memo {
		sort.Strings(v)
	}
	out := make([]string, 0, len(tickets)+1)

	var dfs func(begin string)

	dfs = func(begin string) {
		for len(memo[begin]) > 0 {
			next := memo[begin][0]
			memo[begin] = memo[begin][1:]

			dfs(next)
		}
		out = append(out, begin)
	}
	dfs("JFK")
	return reverse(out)
}

func reverse(ss []string) []string {
	for i := len(ss)/2 - 1; i >= 0; i-- {
		opp := len(ss) - 1 - i
		ss[i], ss[opp] = ss[opp], ss[i]
	}
	return ss
}
