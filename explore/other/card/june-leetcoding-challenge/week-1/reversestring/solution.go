package reversestring

func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		opp := len(s) - 1 - i
		s[i], s[opp] = s[opp], s[i]
	}
}
