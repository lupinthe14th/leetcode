package climbingstairs

func climbStairs(n int) int {
	memo := make(map[int]int)

	var helper func(n int) int
	helper = func(n int) int {
		if n == 0 || n == 1 {
			return 1
		}
		num, ok := memo[n]
		if ok {
			return num
		}
		num = helper(n-1) + helper(n-2)
		memo[n] = num
		return num
	}
	return helper(n)
}
