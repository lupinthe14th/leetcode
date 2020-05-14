package implementtrie

const ALPABET_SIZE = 26

type Trie struct {
	root *Node
}

type Node struct {
	children [ALPABET_SIZE]*Node
	eow      bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{new(Node)}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	probe := this.root
	for i := 0; i < len(word); i++ {
		if probe.children[word[i]-'a'] == nil {
			probe.children[word[i]-'a'] = new(Node)
		}
		probe = probe.children[word[i]-'a']
	}
	probe.eow = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	probe := this.searchPrefix(word)
	return probe != nil && probe.eow
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	return this.searchPrefix(prefix) != nil
}

func (this *Trie) searchPrefix(word string) *Node {
	probe := this.root
	for i := 0; i < len(word); i++ {
		if probe.children[word[i]-'a'] == nil {
			return nil
		}
		probe = probe.children[word[i]-'a']
	}
	return probe
}
