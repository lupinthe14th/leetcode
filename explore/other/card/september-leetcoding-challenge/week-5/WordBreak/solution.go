package wordbreak

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	memo := make(map[string]bool)
	for _, word := range wordDict {
		memo[word] = true
	}
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if memo[s[j:i]] && dp[j] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}
