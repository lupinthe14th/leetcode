package firstuniquecharacterinastring

func firstUniqChar(s string) int {
	memo := make([]int, 256)
	sb := []byte(s)
	l := len(sb)
	for i := 0; i < l; i++ {
		memo[sb[i]]++
	}

	for i := 0; i < l; i++ {
		if memo[sb[i]] == 1 {
			return i
		}
	}
	return -1
}
