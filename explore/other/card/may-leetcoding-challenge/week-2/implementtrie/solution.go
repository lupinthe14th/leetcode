package implementtrie

import "strings"

type Trie struct {
	m map[string]bool
	s []string
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{m: make(map[string]bool), s: make([]string, 0)}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	this.m[word] = true
	this.s = append(this.s, word)
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	return this.m[word]
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for _, word := range this.s {
		if strings.HasPrefix(word, prefix) {
			return true
		}
	}
	return false
}
