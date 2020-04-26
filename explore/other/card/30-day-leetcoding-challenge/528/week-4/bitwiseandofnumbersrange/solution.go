package bitwiseandofnumbersrange

// SeeAlso:
// - https://leetcode.com/problems/bitwise-and-of-numbers-range/discuss/56719/JavaPython-easy-solution-with-explanation
func rangeBitwiseAnd(m int, n int) int {
	var i uint
	for i = 0; m != n; i++ {
		m >>= 1
		n >>= 1
	}
	return n << i
}
