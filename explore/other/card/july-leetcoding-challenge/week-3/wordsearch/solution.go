package wordsearch

type TrieNode struct {
	word     string
	children [58]*TrieNode
}

func exist(board [][]byte, word string) bool {
	out := ""
	root := &TrieNode{}

	node := root
	for _, c := range word {
		if node.children[c-'A'] == nil {
			node.children[c-'A'] = &TrieNode{}
		}
		node = node.children[c-'A']
	}
	node.word = word

	var dfs func(i, j int, node *TrieNode)

	dfs = func(i, j int, node *TrieNode) {
		if i < 0 || j < 0 || i == len(board) || j == len(board[0]) {
			return
		}

		c := board[i][j]

		if c == '#' || node.children[c-'A'] == nil {
			return
		}

		node = node.children[c-'A']

		if node.word != "" {
			out = node.word
			node.word = ""
		}

		board[i][j] = '#'
		dfs(i+1, j, node)
		dfs(i-1, j, node)
		dfs(i, j+1, node)
		dfs(i, j-1, node)
		board[i][j] = c

	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			dfs(i, j, root)
		}
	}
	return out == word
}
