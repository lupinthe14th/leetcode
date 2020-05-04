package ransomnote

func canConstruct(ransomNote string, magazine string) bool {
	memo := make(map[rune]int)

	for _, r := range magazine {
		memo[r]++
	}

	for _, r := range ransomNote {
		c, ok := memo[r]
		if !ok {
			return false
		}
		if c <= 0 {
			return false
		}
		memo[r]--
	}
	return true
}
