package main

func removePalindromeSub(s string) int {
	var ans int
	if len(s) == 0 {
		ans = 0
	} else if isPalindrome(s) {
		ans = 1
	} else {
		ans = 2
	}
	return ans
}

func isPalindrome(s string) bool {
	l := len(s)
	for i := 0; i < l; i++ {
		if s[i] != s[l-1-i] {
			return false
		}
	}
	return true
}

func main() {}
