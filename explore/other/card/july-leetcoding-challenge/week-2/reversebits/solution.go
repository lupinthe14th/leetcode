package reversebits

// Approach 1: Bit by Bit
func reverseBits(num uint32) uint32 {
	var out uint32
	var pow uint32 = 31
	for num != 0 {
		out += num & 1 << pow
		num = num >> 1
		pow -= 1
	}
	return out
}
