package longestcommonsubsequence

func longestCommonSubsequence(text1, text2 string) int {
	m, n := len(text1), len(text2)
	c := make([][]int, m+1)
	for i := range c {
		c[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				c[i][j] = c[i-1][j-1] + 1
			} else {
				c[i][j] = max(c[i-1][j], c[i][j-1])
			}
		}
	}
	return c[m][n]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
