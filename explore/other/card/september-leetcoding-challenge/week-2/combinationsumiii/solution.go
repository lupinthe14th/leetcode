package combinationsumiii

// SeeAlso: https://leetcode.com/problems/combination-sum-iii/discuss/550794/go-backtrack
func combinationSum3(k int, n int) [][]int {
	out := make([][]int, 0)

	var backtrack func(k, n, start int, cur []int)
	backtrack = func(k, n, start int, cur []int) {
		if k == 0 && n == 0 {
			out = append(out, append([]int{}, cur...))
			return
		}

		if k == 0 || n == 0 {
			return
		}

		for i := start; i <= 9; i++ {
			backtrack(k-1, n-i, i+1, append(cur, i))
		}
	}

	backtrack(k, n, 1, nil)
	return out
}
