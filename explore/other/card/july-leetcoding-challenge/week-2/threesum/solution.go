package threesum

import "sort"

func threeSum(nums []int) [][]int {
	const target = 0
	out := make([][]int, 0)
	sort.Ints(nums)

	for i, v := range nums {
		if i > 0 && v == nums[i-1] {
			continue
		}

		lo, hi := i+1, len(nums)-1
		for lo < hi {
			sum := v + nums[lo] + nums[hi]
			if sum > target {
				hi--
			} else if sum < target {
				lo++
			} else {
				out = append(out, []int{v, nums[lo], nums[hi]})
				lo++
				for nums[lo] == nums[lo-1] && lo < hi {
					lo++
				}
			}
		}
	}
	return out
}
