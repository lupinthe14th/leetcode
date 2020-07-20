package addbinary

func addBinary(a string, b string) string {

	l := max(len(a), len(b))
	r := make([]byte, l)
	a = fixCol(a, l)
	b = fixCol(b, l)
	t := false
	for i := l - 1; i >= 0; i-- {
		switch {
		case a[i] == '1' && b[i] == '1':
			if t {
				r[i] = '1'
			} else {
				r[i] = '0'
			}
			t = true
		case a[i] == '0' && b[i] == '0':
			if t {
				r[i] = '1'
			} else {
				r[i] = '0'
			}
			t = false
		default:
			if t {
				r[i] = '0'
				t = true
			} else {
				r[i] = '1'
				t = false
			}
		}
	}
	if t {
		r = append([]byte{'1'}, r...)
	}
	return string(r)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func fixCol(s string, l int) string {
	if len(s) == l {
		return s
	}

	b := make([]byte, l)
	for i := 0; i < l; i++ {
		if i < l-len(s) {
			b[i] = '0'
		} else {
			b[i] = s[i-l+len(s)]
		}
	}
	return string(b)
}
