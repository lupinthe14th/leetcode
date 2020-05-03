package laststoneweight

import (
	"sort"
)

func lastStoneWeight(stones []int) int {
	for len(stones) > 1 {
		sort.Ints(stones)
		x := stones[len(stones)-2]
		y := stones[len(stones)-1]
		stones = stones[:len(stones)-2]
		if x != y {
			stones = append(stones, y-x)
		}
	}
	if len(stones) == 0 {
		return 0
	}
	return stones[0]
}
