package validpalindrome

func isPalindrome(s string) bool {
	b := make([]byte, 0, len(s))
	for i := range s {
		if '0' <= s[i] && s[i] <= '9' || 'a' <= s[i] && s[i] <= 'z' || 'A' <= s[i] && s[i] <= 'Z' {
			b = append(b, s[i]|' ')
		}
	}
	for l, r := 0, len(b)-1; l < r; l, r = l+1, r-1 {
		if b[l] != b[r] {
			return false
		}
	}
	return true
}
