package besttimetobuyandsellstockii

func maxProfit(prices []int) int {
	d := len(prices)
	stock, profit := -1, 0
	for i := 0; i < d; i++ {
		if i < d-1 {
			switch {
			case stock == -1 && prices[i+1]-prices[i] > 0:
				stock = prices[i]
			case stock != -1 && prices[i+1]-prices[i] < 0:
				profit += prices[i] - stock
				stock = -1
			}
		} else {
			if stock != -1 && prices[i]-stock > 0 {
				profit += prices[i] - stock
				stock = -1
			}
		}
	}
	return profit
}
