package adddigits

func addDigits(num int) int {
	if num < 10 {
		return num
	}
	n := num % 9
	if n == 0 {
		return 9
	}
	return n
}
