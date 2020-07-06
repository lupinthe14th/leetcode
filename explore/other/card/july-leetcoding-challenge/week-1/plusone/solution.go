package plusone

func plusOne(digits []int) []int {
	l := len(digits)

	t := 1
	for i := l - 1; i >= 0; i-- {
		t = digits[i] + t
		if t > 9 {
			digits[i] = t % 10
			t /= 10
			if i-1 < 0 {
				digits = append([]int{t}, digits...)
			}
		} else {
			digits[i] = t
			break
		}

	}
	return digits
}
