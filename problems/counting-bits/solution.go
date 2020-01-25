package main

func countBits(num int) []int {
	ans := make([]int, num+1)
	ans[0] = 0
	for i := 1; i < num+1; i++ {
		ans[i] = ans[i&(i-1)] + 1
	}
	return ans
}

func main() {}
