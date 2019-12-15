package sequentialdigits

import (
	"strconv"
)

func sequentialDigits(low int, high int) []int {

	var digit, maxdigit, n int
	var ans []int
	l, h := low, high
	for l > 0 {
		n = l % 10
		l /= 10
		digit++

	}
	for h > 0 {
		h /= 10
		maxdigit++
	}
	for d := digit; d <= maxdigit; d++ {
		for i := n; i < 9; i++ {
			var tmp string
			tmp += strconv.Itoa(i)
			for j := 0; j < d-1; j++ {
				tmp += strconv.Itoa(i + (j + 1))
			}
			c, _ := strconv.Atoi(tmp)
			if c >= low && c <= high {
				ans = append(ans, c)
			} else {
				break
			}
		}
	}
	return ans
}
