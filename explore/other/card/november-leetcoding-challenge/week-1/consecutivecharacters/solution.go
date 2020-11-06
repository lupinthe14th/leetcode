package consecutivecharacters

func maxPower(s string) int {

	out := 1
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	tmp := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			tmp++
		} else {
			tmp = 1
		}
		out = max(out, tmp)
	}
	return out
}
