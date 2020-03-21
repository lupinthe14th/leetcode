package findthelongestsubstringcontainingvowelsinevencounts

func findTheLongestSubstring(s string) int {

	ans := 0
	for i := 0; i < len(s); i++ {
		for j := len(s); j > i; j-- {
			vowels := make(map[rune]int, 5)
			vowels['a'] = 0
			vowels['e'] = 0
			vowels['i'] = 0
			vowels['o'] = 0
			vowels['u'] = 0

			for _, r := range s[i:j] {
				if r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' {
					vowels[r]++
				}
			}
			if count(vowels) {
				ans = max(ans, len(s[i:j]))
			}
		}
	}
	return ans
}

func count(m map[rune]int) bool {
	if m['a']%2 == 0 && m['e']%2 == 0 && m['i']%2 == 0 && m['o']%2 == 0 && m['u']%2 == 0 {
		return true
	}
	return false
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
