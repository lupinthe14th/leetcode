package main

// canJump is determined if you are able to reach the last index.
// Algorthmus: Greedy
func canJump(nums []int) bool {
	// Iterating right-to-left, for each position we check if there is a potential jump that reaches a GOOD index.
	l := len(nums)
	gp := l - 1
	for i := l - 1; i >= 0; i-- {
		if nums[i]+i >= gp {
			gp = i
		}

	}
	return gp == 0
}

func main() {}
