package sequentialdigits

import (
	"strconv"
)

func sequentialDigits(low int, high int) []int {

	var digit, maxdigit int
	var ans []int
	l, h := low, high
	for l > 0 {
		l /= 10
		digit++

	}
	for h > 0 {
		h /= 10
		maxdigit++
	}
	for d := digit; d <= maxdigit; d++ {
		for i := 1; i < 9; i++ {
			var tmp string
			tmp += strconv.Itoa(i)
			for j := 1; j < d; j++ {
				tmp += strconv.Itoa(i + j)
			}
			c, _ := strconv.Atoi(tmp)
			if len(tmp) == d && c >= low && c <= high {
				ans = append(ans, c)
			}
		}
	}
	return ans
}
