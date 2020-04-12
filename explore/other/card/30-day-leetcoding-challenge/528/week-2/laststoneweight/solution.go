package laststoneweight

import (
	"sort"
)

func lastStoneWeight(stones []int) int {
	for len(stones) > 1 {
		sort.Ints(stones)
		x := stones[len(stones)-2]
		y := stones[len(stones)-1]
		if x == y {
			//2 delete
			stones = stones[:len(stones)-2]
		}
		if x != y {
			//replace
			stones[len(stones)-2] = y - x
			//1 delete
			stones = stones[:len(stones)-1]
		}
		if len(stones) == 0 {
			return 0

		}
	}
	return stones[0]
}
