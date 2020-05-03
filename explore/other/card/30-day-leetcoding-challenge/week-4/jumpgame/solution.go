package jumpgame

type Index int

const (
	GOOD Index = iota
	BAD
	UNKNOWN
)

func canJump(nums []int) bool {

	memo := make([]Index, len(nums))
	l := len(nums)

	for i := range nums {
		memo[i] = UNKNOWN
	}

	memo[l-1] = GOOD
	return canJumpFromPosition(0, nums, memo)
}

func canJumpFromPosition(p int, nums []int, memo []Index) bool {
	if memo[p] != UNKNOWN {
		return memo[p] == GOOD
	}

	fJ := min(p+nums[p], len(nums)-1)

	for nj := p + 1; nj <= fJ; nj++ {
		if canJumpFromPosition(nj, nums, memo) {
			memo[p] = GOOD
			return true
		}
	}

	memo[p] = BAD
	return false
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
