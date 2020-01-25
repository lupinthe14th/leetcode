package main

func reverseBits(num uint32) uint32 {
	var ans uint32
	for i := 0; i < 32; i++ {
		if num>>uint(i)&1 == 1 {
			ans |= 1 << uint(31-i)
		}
	}
	return ans
}

func main() {}
