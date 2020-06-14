package largestdivisiblesubset

import "sort"

// https://leetcode.com/problems/largest-divisible-subset/discuss/84006/Classic-DP-solution-similar-to-LIS-O(n2)
func largestDivisibleSubset(nums []int) []int {
	l := len(nums)
	count := make([]int, l)
	pre := make([]int, l)

	sort.Ints(nums)
	max, idx := 0, -1

	for i := 0; i < l; i++ {
		count[i] = 1
		pre[i] = -1

		for j := i - 1; j >= 0; j-- {
			if nums[i]%nums[j] == 0 {
				if 1+count[j] > count[i] {
					count[i] = count[j] + 1
					pre[i] = j
				}
			}
		}
		if count[i] > max {
			max = count[i]
			idx = i
		}
	}
	out := make([]int, 0)
	for idx != -1 {
		out = append(out, nums[idx])
		idx = pre[idx]
	}
	sort.Ints(out)
	return out
}
