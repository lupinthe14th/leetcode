package threesum

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	memo := make(map[int]int)
	for i, n := range nums {
		memo[n] = i
	}
	seen := make(map[[3]int]bool)
	out := make([][]int, 0)
	for i, x := range nums {
		for j := i + 1; j < len(nums); j++ {
			y := nums[j]
			if k, ok := memo[-(y + x)]; ok && i != k && j != k && k > j {
				t := []int{nums[i], nums[j], nums[k]}
				sort.Ints(t)
				var tmp [3]int
				for i := range t {
					tmp[i] = t[i]
				}
				if !seen[tmp] {
					out = append(out, t)
					seen[tmp] = true
				}
			}
		}
	}
	return out
}
