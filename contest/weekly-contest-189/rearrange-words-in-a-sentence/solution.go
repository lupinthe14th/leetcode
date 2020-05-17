package rearrangewordsinasentence

import (
	"sort"
	"strings"
)

func arrangeWords(text string) string {

	texts := strings.Split(strings.ToLower(text), " ")

	sort.SliceStable(texts, func(i, j int) bool {
		return len(texts[i]) < len(texts[j])
	})
	var b []byte
	for i, r := range texts[0] {
		if i == 0 {
			r = (r - 'a') + 'A'
		}
		b = append(b, byte(r))
	}
	texts[0] = string(b)

	return strings.Join(texts, " ")
}
