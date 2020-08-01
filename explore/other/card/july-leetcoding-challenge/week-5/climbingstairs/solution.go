package climbingstairs

func climbStairs(n int) int {
	memo := make([]int, n+1)

	var helper func(n int) int
	helper = func(n int) int {
		if n == 0 || n == 1 {
			return 1
		}
		if memo[n] > 0 {
			return memo[n]
		}
		memo[n] = helper(n-1) + helper(n-2)
		return memo[n]
	}
	return helper(n)
}
