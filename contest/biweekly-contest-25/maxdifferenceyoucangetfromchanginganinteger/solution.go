package maxdifferenceyoucangetfromchanginganinteger

import (
	"strconv"
	"strings"
)

// https://leetcode.com/problems/max-difference-you-can-get-from-changing-an-integer/discuss/608703/Python3-Easy-Solution
func maxDiff(num int) int {
	n := strconv.Itoa(num)
	maxNum, minNum := -1<<31, 1<<31

	for _, i := range "0123456789" {
		for _, j := range "0123456789" {
			nextNum := strings.ReplaceAll(n, string(i), string(j))
			t, _ := strconv.Atoi(nextNum)
			if nextNum[0] == '0' || t == 0 {
				continue
			}
			maxNum = max(maxNum, t)
			minNum = min(minNum, t)
		}
	}
	return maxNum - minNum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
