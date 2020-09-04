package repeatedsubstringpattern

func repeatedSubstringPattern(s string) bool {
	l := len(s)
	for i := l / 2; i > 0; i-- {
		tmp := ""
		for len(tmp) < l {
			tmp += s[:i]
		}
		if tmp == s {
			return true
		}
	}
	return false
}
