package besttimetobuyandsellstock

func maxProfit(price []int) int {
	// 一番安い値を探す。
	// 安い値に更新されたら次の日（i）に移動する。
	// 安い値に更新されなかったら、最大利益を計算する。
	// 利益が前の利益より大きければ利益を更新する。
	// それを最後の日まで繰り返せば、最終的に最大の利益が得られる。
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
