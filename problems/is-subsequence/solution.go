package issubsequence

func isSubsequence(s string, t string) bool {
	dp := make([]int, len(s)+1)
	for i := 0; i < len(s)+1; i++ {
		dp[i] = -1
	}
	for i := 1; i <= len(s); i++ {
		j := 0
		if dp[i-1] != -1 {
			j = dp[i-1] + 1
		}
		for j < len(t) {
			if s[i-1] == t[j] {
				dp[i] = j
				break
			}
			j++
		}
		if dp[i] == -1 {
			return false
		}
	}
	return true
}
