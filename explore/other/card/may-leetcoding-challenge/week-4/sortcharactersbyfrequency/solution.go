package sortcharactersbyfrequency

func frequencySort(s string) string {
	maxl := 0
	out := make([]byte, 0, len(s))

	memo := make([]int, 256)
	for _, b := range s {
		memo[b]++
		maxl = max(maxl, memo[b])
	}

	for maxl > 0 {
		for b, c := range memo {
			if c == maxl {
				for i := 0; i < c; i++ {
					out = append(out, byte(b))
				}
			}
		}
		maxl--
	}
	return string(out)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
