package besttimetobuyandsellstock

func maxProfit(price []int) int {
	var a int

	for i := 0; i < len(price); i++ {
		for j := i + 1; j < len(price); j++ {
			profit := price[j] - price[i]
			if profit > 0 {
				a = max(a, profit)
			}
		}

	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
