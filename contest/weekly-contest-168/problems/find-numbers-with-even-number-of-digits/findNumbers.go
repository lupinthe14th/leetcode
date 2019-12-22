package findnumberswithevennumberofdigits

func findNumbers(nums []int) int {
	var c int
	for _, v := range nums {
		if numberOfDigits(v)&1 == 0 {
			c++
		}
	}
	return c
}

func numberOfDigits(n int) int {
	var c int
	if n == 0 {
		return 1
	}
	for n > 0 {
		n /= 10
		c++
	}
	return c
}
