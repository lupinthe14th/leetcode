package minimumnumberofstepstomaketwostringsanagram

func minSteps(s string, t string) int {
	l := len(s)
	m := make(map[rune]int, l)

	for _, v := range t {
		m[v]++
	}

	for _, v := range s {
		if _, ok := m[v]; ok {
			if m[v] > 0 {
				m[v]--
			}
		}
	}

	var c int
	for _, v := range m {
		c += v
	}
	return c
}
