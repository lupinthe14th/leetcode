package constructkpalindromestrings

func canConstruct(s string, k int) bool {
	if len(s) < k {
		return false
	}
	dp := make([]int, 26)
	for _, r := range s {
		dp[r-'a']++
	}
	odd := 0
	for _, n := range dp {
		if n%2 != 0 {
			odd++
		}
	}
	return odd <= k
}
