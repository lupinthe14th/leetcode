package increasingdecreasingstring

func sortString(s string) string {
	m := make(map[rune]int, len(s))
	for _, r := range s {
		m[r]++
	}

	ans := make([]rune, 0, len(s))
	var i int = len(s)
	for i > 0 {
		for j := 97; j < 123; j++ {
			if n := m[rune(j)]; n > 0 {
				ans = append(ans, rune(j))
				m[rune(j)]--
				i--
			}
			if i < 0 {
				break
			}
		}
		for j := 122; j > 96; j-- {
			if n := m[rune(j)]; n > 0 {
				ans = append(ans, rune(j))
				m[rune(j)]--
				i--
			}
			if i < 0 {
				break
			}
		}
	}
	return string(ans)
}
