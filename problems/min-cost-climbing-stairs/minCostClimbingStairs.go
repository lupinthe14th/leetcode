package mincostclimbingstairs

func minCostClimbingStairs(cost []int) int {
	l := len(cost)
	memo := make([]int, l+1)

	for i := 2; i <= l; i++ {
		memo[i] = min(memo[i-1]+cost[i-1], memo[i-2]+cost[i-2])
	}
	return memo[l]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
