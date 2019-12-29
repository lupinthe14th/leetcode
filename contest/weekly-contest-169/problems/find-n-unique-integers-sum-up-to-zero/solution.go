package main

func sumZero(n int) []int {
	ans := make([]int, n)

	// n is odd
	if n&1 == 1 {
		mid := n / 2
		for i := 1; i <= n/2; i++ {
			ans[mid-i] = -1 * i
			ans[mid+i] = i
		}
	} else {
		mid := n / 2
		for i := 1; i <= n/2; i++ {
			ans[mid-i] = -i
			ans[mid+i-1] = i
		}
	}
	return ans
}

func main() {}
