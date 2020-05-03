package happynumber

func isHappy(n int) bool {
	memo := make(map[int]bool)

	for n != 1 {
		n = sumOfSquaresOfDigits(n)
		if memo[n] {
			return false
		}
		memo[n] = true
	}
	return true
}

func sumOfSquaresOfDigits(n int) int {
	var sum int
	for n > 0 {
		d := n % 10
		sum += d * d
		n /= 10
	}
	return sum
}
