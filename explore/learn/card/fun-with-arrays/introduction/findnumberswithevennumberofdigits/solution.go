package findnumberswithevennumberofdigits

func findNumbers(nums []int) int {

	var digitNumber = func(n int) int {
		out := 0
		for n > 0 {
			n /= 10
			out++
		}
		return out
	}
	out := 0
	for _, num := range nums {
		if digitNumber(num)&1 == 0 {
			out++
		}
	}
	return out
}
