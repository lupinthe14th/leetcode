package reformatthestring

func reformat(s string) string {
	l := len(s)
	if l == 1 {
		return s
	}
	chars := make([]rune, 0, l/2+1)
	digits := make([]rune, 0, l/2+1)

	for _, r := range s {
		switch {
		case r >= 'a' && r <= 'z':
			chars = append(chars, r)
		case r >= '0' && r <= '9':
			digits = append(digits, r)
		}
	}
	cl := len(chars)
	dl := len(digits)
	if cl == 0 || dl == 0 {
		return ""
	}
	if cl > dl && cl-1 != dl {
		return ""
	}
	if dl > cl && dl-1 != cl {
		return ""
	}
	ans := make([]rune, l)
	j, k := 0, 0
	if cl > dl {
		for i := range ans {
			if i&1 == 0 {
				ans[i] = chars[j]
				j++
			} else {
				ans[i] = digits[k]
				k++
			}
		}
	} else {
		for i := range ans {
			if i&1 == 1 {
				ans[i] = chars[j]
				j++
			} else {
				ans[i] = digits[k]
				k++
			}
		}

	}
	return string(ans)
}
