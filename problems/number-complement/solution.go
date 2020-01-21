package main

import (
	"math"
)

func findComplement(num int) int {
	var ans int
	l := int(math.Log2(float64(num))) + 1
	for i := 0; i < l; i++ {
		if num>>uint(i)&1 == 1 {
			ans &^= 1 << uint(i)
		} else {
			ans |= 1 << uint(i)
		}
	}
	return ans
}
func main() {}
