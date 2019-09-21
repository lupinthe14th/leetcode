package maxnumberofapples

import (
	"sort"
)

func maxNumberOfApples(arr []int) int {
	var sum int
	sort.Ints(arr)
	for i, w := range arr {
		sum = sum + w
		if sum > 5000 {
			return i
		}
	}
	return len(arr)
}
