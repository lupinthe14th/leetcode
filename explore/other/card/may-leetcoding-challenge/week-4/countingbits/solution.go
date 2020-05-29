package countingbits

func countBits(num int) []int {
	out := []int{0, 1}
	if num == 0 {
		return []int{0}
	}
	if num == 1 {
		return out
	}
	b, c := 1, 0
	for i := 2; i <= num; i++ {
		if i == 2<<uint(b) {
			c = 0
			out = append(out, out[c]+1)
			b++
			c++
		} else {
			out = append(out, out[c]+1)
			c++
		}
	}
	return out
}
