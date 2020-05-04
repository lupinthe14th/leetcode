package numbercomplement

import (
	"math"
	"math/bits"
)

// https://leetcode.com/problems/number-complement/discuss/488055/Python-O(-lg-n-)-sol.-by-XOR-masking.-85+-With-explanation
func findComplement(num int) int {
	bitMask := math.Pow(2, float64(bits.Len32(uint32(num)))) - 1
	return num ^ int(bitMask)
}
