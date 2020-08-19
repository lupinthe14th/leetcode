package numberswithsameconsecutivedifferences

func numsSameConsecDiff(N int, K int) []int {
	out := make([]int, 0)

	var helper func(n int, nums []int)

	helper = func(n int, nums []int) {
		if len(nums) == N {
			if len(nums) != 1 && nums[0] == 0 {
				return
			}
			num, e := 0, 1
			for i := N - 1; i >= 0; i-- {
				num += nums[i] * e
				e *= 10
			}
			out = append(out, num)
			return
		}
		if K == 0 {
			helper(n, append(nums, n))
			return
		}
		if n+K <= 9 {
			helper(n+K, append(nums, n+K))
		}
		if n-K >= 0 {
			helper(n-K, append(nums, n-K))
		}
	}
	for i := 0; i <= 9; i++ {
		num := make([]int, 0, N)
		num = append(num, i)
		helper(i, num)
	}
	return out
}
