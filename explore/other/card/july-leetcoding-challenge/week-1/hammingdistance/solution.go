package hammingdistance

func hammingDistance(x int, y int) int {
	b := x ^ y
	c := 0
	for b > 0 {
		if b&1 == 1 {
			c++
		}
		b = b >> 1
	}
	return c
}
