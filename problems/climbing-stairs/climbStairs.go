package climbingstairs

func climbStairs(n int) int {

	memo := make(map[int]int, n)

	return climbingstairs(n, memo)
}

func climbingstairs(n int, memo map[int]int) int {
	if n == 0 || n == 1 {
		return 1
	}
	if _, ok := memo[n]; !ok {
		memo[n] = climbingstairs(n-1, memo) + climbingstairs(n-2, memo)
	}
	return memo[n]
}
