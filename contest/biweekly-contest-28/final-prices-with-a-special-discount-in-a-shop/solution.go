package finalpriceswithaspecialdiscountinashop

func finalPrices(prices []int) []int {
	out := make([]int, 0, len(prices))
	for i := 0; i < len(prices); i++ {
		p := prices[i]
		for j := i + 1; j < len(prices); j++ {
			if prices[j] <= p {
				p = p - prices[j]
				break
			}
		}
		out = append(out, p)
	}
	return out
}
