package squaresofasortedarray

import "sort"

func sortedSquares(nums []int) []int {
	out := make([]int, len(nums))
	for i, num := range nums {
		out[i] = num * num
	}
	sort.Ints(out)
	return out
}
