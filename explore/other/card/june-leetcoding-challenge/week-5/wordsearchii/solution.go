package wordsearchii

import "sort"

// https://leetcode.com/problems/word-search-ii/discuss/542405/Go-DFS-+-Trie-solution
type TrieNode struct {
	word     string
	children [26]*TrieNode
}

func findWords(board [][]byte, words []string) []string {
	out := make([]string, 0)
	root := &TrieNode{}

	for _, w := range words {
		node := root
		for _, c := range w {
			if node.children[c-'a'] == nil {
				node.children[c-'a'] = &TrieNode{}
			}
			node = node.children[c-'a']
		}
		node.word = w
	}

	var dfs func(i, j int, node *TrieNode)
	dfs = func(i, j int, node *TrieNode) {
		if i < 0 || j < 0 || i == len(board) || j == len(board[0]) {
			return
		}

		c := board[i][j]

		if c == '#' || node.children[c-'a'] == nil {
			return
		}

		node = node.children[c-'a']
		if node.word != "" {
			out = append(out, node.word)
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
		for j := 0; j < len(board[i]); j++ {
			dfs(i, j, root)
		}
	}
	sort.Strings(out)
	return out
}
