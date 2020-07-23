package singlenumberiii

func singleNumber(nums []int) []int {
	// Pass 1:
	// Get the XOR of the two numbers we need to find
	diff := 0
	for _, n := range nums {
		diff ^= n
	}
	// Get its last set bit
	diff &= -diff

	// Pass 2:
	out := make([]int, 2)
	for _, n := range nums {
		if n&diff == 0 {
			// the bit is not set
			out[0] ^= n
		} else {
			// the bit is set
			out[1] ^= n
		}
	}
	return out
}
