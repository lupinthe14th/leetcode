package main

import "strconv"

func getNoZeroIntegers(n int) []int {
	ans := make([]int, 2)
	for a := 1; a < n; a++ {
		b := n - a

		if !isZeroInteger(a) && !isZeroInteger(b) {
			ans[0] = a
			ans[1] = b
			break
		}
	}

	return ans
}

func isZeroInteger(x int) bool {
	s := strconv.Itoa(x)
	for _, v := range s {
		if v == '0' {
			return true
		}
	}
	return false
}

func main() {}
