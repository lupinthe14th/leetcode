package firstuniquecharacterinastring

func firstUniqChar(s string) int {
	memo := make(map[rune]int, 26)
	for _, r := range s {
		memo[r]++
	}

	for k, v := range memo {
		if v != 1 {
			delete(memo, k)
		}
	}

	if len(memo) != 0 {
		for i, r := range s {
			if _, ok := memo[r]; ok {
				return i
			}
		}
	}
	return -1
}
