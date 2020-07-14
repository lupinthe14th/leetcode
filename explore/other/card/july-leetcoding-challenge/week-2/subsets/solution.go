package subsets

// Approach 2: Backtracking
func subsets(nums []int) [][]int {
	out := make([][]int, 0)

	var backtrack func(nums, cur []int)

	backtrack = func(nums, cur []int) {
		out = append(out, append([]int{}, cur...))
		for i := 0; i < len(nums); i++ {
			backtrack(nums[i+1:], append(cur, nums[i]))
		}
	}

	backtrack(nums, nil)
	return out
}
