package climbingstairs

import "math/big"

func climbStairs(n int) *big.Int {

	memo := make(map[int]*big.Int, n)

	return climbingstairs(n, memo)
}

func climbingstairs(n int, memo map[int]*big.Int) *big.Int {
	if n < 0 {
		return big.NewInt(0)
	}
	if n == 0 || n == 1 {
		return big.NewInt(1)
	}
	var ans big.Int
	if _, ok := memo[n]; !ok {
		ans.Add(climbingstairs(n-1, memo), climbingstairs(n-2, memo))
		memo[n] = &ans
	}
	return memo[n]
}
