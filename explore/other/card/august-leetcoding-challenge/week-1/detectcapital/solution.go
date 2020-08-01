package detectcapital

func detectCapitalUse(word string) bool {
	isUpper := func(s string) bool {
		for _, r := range s {
			if 'A' <= r && r <= 'Z' {
				continue
			}
			return false
		}
		return true
	}
	isLower := func(s string) bool {
		for _, r := range s {
			if 'a' <= r && r <= 'z' {
				continue
			}
			return false
		}
		return true
	}

	switch {
	case isUpper(word):
		return true
	case isLower(word):
		return true
	case isUpper(string(word[0])) && isLower(string(word[1:])):
		return true
	}
	return false
}
