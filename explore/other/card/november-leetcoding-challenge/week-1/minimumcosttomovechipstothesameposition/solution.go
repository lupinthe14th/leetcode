package minimumcosttomovechipstothesameposition

func minCostToMoveChips(position []int) int {
	o, e := 0, 0
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	for _, v := range position {
		if v%2 == 0 {
			e++
		} else {
			o++
		}
	}
	return min(o, e)
}
