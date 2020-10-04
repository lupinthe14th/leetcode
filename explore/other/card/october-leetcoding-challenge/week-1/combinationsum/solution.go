package combinationsum

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	out := [][]int{}
	var helper func(v []int, target, idx int)
	helper = func(v []int, target, idx int) {
		if target == 0 {
			out = append(out, v)
			return
		}

		if target < 0 {
			return
		}

		for i := idx; i < len(candidates); i++ {
			if candidates[i] <= target {
				match := make([]int, len(v)+1)
				copy(match, v)
				match[len(v)] = candidates[i]
				helper(match, target-candidates[i], i)
			} else {
				break
			}
		}
	}

	sort.Ints(candidates)
	helper([]int{}, target, 0)

	return out
}
