package findallduplicatesinanarray

import "sort"

func findDuplicates(nums []int) []int {
	out := make([]int, 0)
	if len(nums) == 0 {
		return out
	}
	sort.Ints(nums)
	prev := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == prev {
			out = append(out, nums[i])
		} else {
			prev = nums[i]
		}
	}
	return out
}
