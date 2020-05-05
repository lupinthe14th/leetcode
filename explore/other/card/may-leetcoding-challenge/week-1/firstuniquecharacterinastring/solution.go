package firstuniquecharacterinastring

func firstUniqChar(s string) int {
	memo := make([]int, 256)
	for _, r := range s {
		memo[r]++
	}

	for i, r := range s {
		if memo[r] == 1 {
			return i
		}
	}
	return -1
}
