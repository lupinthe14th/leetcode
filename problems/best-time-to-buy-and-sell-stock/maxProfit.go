package besttimetobuyandsellstock

func maxProfit(price []int) int {
	// 1. Find the lowest price.（一番安い値を探す。）
	// 2. Move to the next day after updating the min parameter to the lowest price.（安い値に更新されたら次の日に移動する。）
	// 3. If the minimum price is not updated, calculate the maximum profit.（安い値に更新されなかったら、最大利益を計算する。）
	// 4. If the profit has increased, update the profit.（利益が前の利益より大きければ利益を更新する。）
	// 5. If you repeat it until the last day, you will ultimately get the most profit.（それを最後の日まで繰り返せば、最終的に最大の利益が得られる。）
	const MaxUint = ^uint(0)
	const MaxInt = int(MaxUint >> 1)

	min := MaxInt
	profit := 0
	for _, p := range price {
		if p < min {
			min = p
			continue
		}

		if p-min > profit {
			profit = p - min
		}
	}
	return profit
}
