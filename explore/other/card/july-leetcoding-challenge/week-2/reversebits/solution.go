package reversebits

// Approach 2: Byte by Byte with Memoization
func reverseBits(num uint32) uint32 {
	var out uint64
	var pow uint64 = 24
	cache := make(map[uint32]uint64)

	reverseByte := func(b uint32) uint64 {
		if v, ok := cache[b]; ok {
			return v
		}

		v := (uint64(b) * 0x0202020202 & 0x010884422010) % 1023
		cache[b] = v
		return v

	}
	for num != 0 {
		out += reverseByte(num&0xff) << pow
		num = num >> 8
		pow -= 8
	}
	return uint32(out)
}
