package main

func decompressRLElist(nums []int) []int {
	ans := make([]int, 0)
	for i := 0; i < len(nums); i += 2 {
		a := nums[i]
		b := nums[i+1]

		for j := 0; j < a; j++ {
			ans = append(ans, b)
		}
	}
	return ans
}

func main() {}
