package wordsearch

func exist(board [][]byte, word string) bool {

	out := false
	memo := make(map[byte][][2]int, 0)

	for i := range board {
		for j := range board[i] {
			k := board[i][j]
			memo[k] = append(memo[k], [2]int{i, j})
		}
	}

	var dfs func(i, j, k int)

	dfs = func(i, j, k int) {
		if i < 0 || j < 0 || i == len(board) || j == len(board[0]) || k == len(word) {
			return
		}

		c := board[i][j]

		if c != word[k] {
			return
		}

		if k == len(word)-1 && c == word[k] {
			out = true
			return
		}

		board[i][j] = '#'
		dfs(i+1, j, k+1)
		dfs(i-1, j, k+1)
		dfs(i, j+1, k+1)
		dfs(i, j-1, k+1)
		board[i][j] = c
	}

	for _, idxes := range memo[word[0]] {
		dfs(idxes[0], idxes[1], 0)
	}
	return out
}
