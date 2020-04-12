package stringmatchinginanarray

func stringMatching(words []string) []string {
	ans := []string{}
	seen := make(map[string]bool, len(words))
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words); j++ {
			if i == j {
				continue
			}
			if seen[words[j]] {
				continue
			}
			if isSubstring(words[i], words[j]) {
				ans = append(ans, words[j])
				seen[words[j]] = true
			}
		}
	}
	return ans
}

func isSubstring(s, ss string) bool {
	for i := 0; i < len(s)-len(ss)+1; i++ {
		if s[i:i+len(ss)] == ss {
			return true
		}
	}
	return false
}
