package playwithchips

func minCostToMoveChips(chips []int) int {

	l := len(chips)
	if l == 1 {
		return 0
	}
	if l%2 == 1 {
		return (l - 1) / 2
	}
	return l / 2

}
