package performstringshifts

func stringShift(s string, shift [][]int) string {
	b := make([]byte, len(s))

	for i := range s {
		b[i] = s[i]
	}
	for i := range shift {
		direction := shift[i][0]
		amount := shift[i][1]
		if direction == 0 {
			// left shift
			b = append(b[amount:], b[:amount]...)
		} else {
			// right shift
			b = append(b[len(b)-amount:], b[:len(b)-amount]...)
		}
	}
	return string(b)
}
