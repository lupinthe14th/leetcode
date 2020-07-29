package besttimetobuyandsellstockwithcooldown

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/discuss/75928/Share-my-DP-solution-(By-State-Machine-Thinking)
func maxProfit(prices []int) int {
	l := len(prices)
	sold, hold, rest := 0, -1<<31, 0
	for i := 0; i < l; i++ {
		prev := sold
		sold = hold + prices[i]
		hold = max(hold, rest-prices[i])
		rest = max(rest, prev)
	}
	return max(sold, rest)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
