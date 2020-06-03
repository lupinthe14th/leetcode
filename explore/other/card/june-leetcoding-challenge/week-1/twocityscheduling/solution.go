package twocityscheduling

import (
	"sort"
)

func twoCitySchedCost(costs [][]int) int {
	out := 0
	l := len(costs)
	type memo struct {
		d, i int
	}
	m := make([]memo, 0, l)
	for i, cost := range costs {
		m = append(m, memo{d: abs(cost[0], cost[1]), i: i})
	}
	sort.Slice(m, func(i, j int) bool {
		return m[i].d > m[j].d
	})

	a, b := 0, 0
	for i := 0; i < l; i++ {
		if a < l/2 && b < l/2 {
			out += min(costs[m[i].i][0], costs[m[i].i][1])
			if costs[m[i].i][0] < costs[m[i].i][1] {
				a++
			} else {
				b++
			}
		} else {
			if a > b {
				out += costs[m[i].i][1]
			} else {
				out += costs[m[i].i][0]
			}
		}
	}
	return out
}

func abs(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
