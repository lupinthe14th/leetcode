package checkifawordoccursasaprefixofanywordinasentence

import "strings"

func isPrefixOfWord(sentence string, searchWord string) int {
	out := -1
	for i, s := range strings.Split(sentence, " ") {
		if strings.HasPrefix(s, searchWord) {
			out = i + 1
			break
		}
	}
	return out
}
