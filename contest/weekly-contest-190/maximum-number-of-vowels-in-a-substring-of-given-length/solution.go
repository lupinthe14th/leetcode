package maximumnumberofvowelsinasubstringofgivenlength

func maxVowels(s string, k int) int {
	var mem [26]int
	for i := 0; i < k; i++ {
		mem[s[i]-'a']++
	}
	out := 0
	for i := 0; i < len(s)-k+1; i++ {
		t := mem[0] + mem[4] + mem[8] + mem[14] + mem[20]
		out = max(out, t)

		if i+k < len(s) {
			mem[s[i]-'a']--
			mem[s[i+k]-'a']++
		}
	}
	return out
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
