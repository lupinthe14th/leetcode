package main

import (
	"fmt"
)

// greedy approach
// 1. 全ての時刻をforで回す。
// 2. ビットの数を調べるため、hを6ビット左にシフトして、h OR mとして、ビットの数を数える。
// 3. ビットの数があっていれば、答えのスライスに時刻フォーマットして追加する
// SeeAlso: https://leetcode.com/problems/binary-watch/discuss/88465/Straight-forward-6-line-c++-solution-no-need-to-explain
func readBinaryWatch(num int) []string {
	ans := make([]string, 0, 10)

	for h := 0; h < 12; h++ {
		for m := 0; m < 60; m++ {
			if bitCount(h<<6|m) == num {
				ans = append(ans, fmt.Sprintf("%0d:%02d", h, m))
			}
		}
	}
	return ans
}

func bitCount(num int) int {
	var c int
	for num > 0 {
		if num&1 == 1 {
			c++
		}
		num >>= 1
	}
	return c
}

func main() {}
