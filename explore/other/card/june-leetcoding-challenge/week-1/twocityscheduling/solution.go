package twocityscheduling

import (
	"sort"
)

func twoCitySchedCost(costs [][]int) int {
	out := 0
	sort.Slice(costs, func(i, j int) bool {
		return costs[i][0]-costs[i][1] < costs[j][0]-costs[j][1]
	})

	for i := 0; i < len(costs)/2; i++ {
		out += costs[i][0] + costs[i+len(costs)/2][1]
	}
	return out
}
