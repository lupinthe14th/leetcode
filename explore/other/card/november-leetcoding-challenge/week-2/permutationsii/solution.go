package permutationsii

import "sort"

func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	sort.Ints(nums)
	out := make([][]int, 0)

	var backtrackUnique func(nums, prev []int)
	backtrackUnique = func(nums, prev []int) {
		if len(nums) == 0 {
			out = append(out, append([]int{}, prev...))
			return
		}

		for i := 0; i < len(nums); i++ {
			if i != 0 && nums[i] == nums[i-1] {
				continue
			}
			backtrackUnique(append(append([]int{}, nums[0:i]...), nums[i+1:]...), append(prev, nums[i]))
		}
	}
	backtrackUnique(nums, nil)
	return out
}
