package nextgreaterelementiii

import (
	"math"
	"sort"
)

func nextGreaterElement(n int) int {
	digits := make([]int, 0)

	for n != 0 {
		digits = append(digits, n%10)
		n /= 10
	}

	for i := 1; i < len(digits); i++ {
		if digits[i] >= digits[i-1] {
			continue
		}
		index, diff := 0, 10
		for j := 0; j < i; j++ {
			curDiff := digits[j] - digits[i]
			if curDiff > 0 && curDiff < diff {
				diff = curDiff
				index = j
			}
		}
		digits[i], digits[index] = digits[index], digits[i]
		sort.Sort(sort.Reverse(sort.IntSlice(digits[:i])))

		dest := 0
		for index, value := range digits {

			dest += value * int(math.Pow10(index))
		}
		if dest <= math.MaxInt32 {
			return dest
		}
	}
	return -1
}
