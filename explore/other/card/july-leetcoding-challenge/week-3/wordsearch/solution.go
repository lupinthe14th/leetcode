package wordsearch

// https://leetcode.com/problems/word-search/discuss/27658/Accepted-very-short-Java-solution.-No-additional-space.
func exist(board [][]byte, word string) bool {
	var dfs func(i, j, k int) bool

	dfs = func(i, j, k int) bool {
		if k == len(word) {
			return true
		}
		if i < 0 || j < 0 || i == len(board) || j == len(board[0]) {
			return false
		}

		c := board[i][j]

		if c == '#' || c != word[k] {
			return false
		}

		board[i][j] = '#'
		out := dfs(i+1, j, k+1) || dfs(i-1, j, k+1) || dfs(i, j+1, k+1) || dfs(i, j-1, k+1)
		board[i][j] = c
		return out
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
