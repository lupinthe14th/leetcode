package main

// Index is the status of each position in the array
type index int

const (
	unknown index = iota
	good
	bad
)

// canJump is determined if you are able to reach the last index.
// Algorthmus: Dynamic Programing Bottom-up
func canJump(nums []int) bool {
	// Initialize memoization table
	l := len(nums)
	memo := make([]index, l)
	for i := range nums {
		memo[i] = unknown
	}
	// 最終地点はGood
	memo[l-1] = good

	// iを最終地点の一つ前から開始点0までループ
	for i := l - 2; i >= 0; i-- {
		fj := min(i+nums[i], l-1)
		for j := i + 1; j <= fj; j++ {
			if memo[j] == good {
				memo[i] = good
				break
			}
		}
	}
	return memo[0] == good
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {}
