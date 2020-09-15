package houserobber

func rob(nums []int) int {

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	l := len(nums)
	dp := make(map[int]int, l)

	var helper func(i int) int
	helper = func(i int) int {
		if i < 0 {
			return 0
		}
		if n, ok := dp[i]; ok {
			return n
		}
		dp[i] = max(helper(i-2)+nums[i], helper(i-1))
		return dp[i]
	}
	return helper(l - 1)
}
