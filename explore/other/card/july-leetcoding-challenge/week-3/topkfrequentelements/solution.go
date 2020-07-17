package topkfrequentelements

import (
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	sort.Ints(nums)
	r := 0
	memo := make(map[int]int)
	for i := range nums {
		memo[nums[i]]++
		r = max(r, memo[nums[i]])
	}

	tmp := make([][]int, r+1)
	for n, v := range memo {
		tmp[v] = append(tmp[v], n)
	}
	out := make([]int, 0, k)
	for i := len(tmp) - 1; i >= 0; i-- {
		out = append(out, tmp[i]...)
		if len(out) >= k {
			break
		}
	}
	return out
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
