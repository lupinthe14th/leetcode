package validpalindrome

func isPalindrome(s string) bool {
	isalnum := func(b byte) bool {
		if '0' <= b && b <= '9' || 'a' <= b && b <= 'z' || 'A' <= b && b <= 'Z' {
			return true
		}
		return false
	}
	l, r := 0, len(s)-1
	for l < r {
		if !isalnum(s[l]) {
			l++
		} else if !isalnum(s[r]) {
			r--
		} else {
			if s[l]|' ' != s[r]|' ' {
				return false
			}
			l++
			r--
		}
	}
	return true
}
