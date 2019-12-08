package subtracttheproductandsumofdigitsofaninteger

func subtractProductAndSum(n int) int {
	var p, s int = 1, 0
	for n > 0 {
		d := n % 10
		p *= d
		s += d
		n /= 10
	}
	return p - s
}
