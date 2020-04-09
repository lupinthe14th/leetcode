package backspacestringcompare

func backspaceCompare(S string, T string) bool {
	return b(S) == b(T)
}

func b(str string) string {
	var ans []rune
	for _, s := range str {
		if len(ans) == 0 && s == '#' {
			continue
		}
		if s == '#' {
			ans = ans[:len(ans)-1]
			continue
		}
		ans = append(ans, s)
	}
	return string(ans)
}
