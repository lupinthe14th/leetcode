package addandsearchworddatastructuredesign

// See: https://leetcode.com/problems/add-and-search-word-data-structure-design/discuss/59554/My-simple-and-clean-Java-code
type WordDictionary struct {
	word     string
	children [26]*WordDictionary
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	for i := range word {
		if this.children[word[i]-'a'] == nil {
			this.children[word[i]-'a'] = &WordDictionary{}
		}
		this = this.children[word[i]-'a']
	}
	this.word = word
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	var match func(b []byte, k int, node *WordDictionary) bool
	match = func(b []byte, k int, node *WordDictionary) bool {
		if k == len(b) {
			return !(node.word == "")
		}
		if b[k] != '.' {
			return node.children[b[k]-'a'] != nil && match(b, k+1, node.children[b[k]-'a'])
		} else {
			for i := 0; i < len(node.children); i++ {
				if node.children[i] != nil && match(b, k+1, node.children[i]) {
					return true
				}
			}
		}
		return false
	}
	b := make([]byte, len(word))
	for i := range word {
		b[i] = word[i]
	}
	return match(b, 0, this)
}
