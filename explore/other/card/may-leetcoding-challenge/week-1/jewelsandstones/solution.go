package jewelsandstones

func numJewelsInStones(J string, S string) int {
	ans := 0
	for i := range J {
		for j := range S {
			if J[i] == S[j] {
				ans++
			}
		}
	}
	return ans
}
