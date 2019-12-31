package main

// Index is the status of each position in the array
type Index int

const (
	Unknown Index = iota
	Good
	Bad
)

// canJump is determined if you are able to reach the last index.
// Algorthmus: Dynamic Programing Top-down
func canJump(nums []int) bool {
	// Initialize memoization table
	l := len(nums)
	memo := make([]Index, l)
	for i := range nums {
		memo[i] = Unknown
	}
	// 最終地点はGood
	memo[l-1] = Good

	return canJumpFromPosition(0, nums, memo)
}

func canJumpFromPosition(p int, nums []int, memo []Index) bool {
	// memoization table check
	if memo[p] != Unknown {
		if memo[p] == Good {
			return true
		}
		return false
	}

	// 最も遠いジャンプ地点を決める
	// スライスにおける現在地点とその値を足した値が最も遠いジャンプ地点。
	// もしそれがスライス長を超えるならスライス長が最も遠いジャンプ地点
	// となるため両者を比較して小さい方を最も遠いジャンプ地点となる。
	fj := min(p+nums[p], len(nums)-1) //  furthest Jump

	// 次の地点を逐次再帰的に調べる。
	// npを最大値から最小値までチェック
	for np := fj; np > p; np-- {
		if canJumpFromPosition(np, nums, memo) {
			memo[p] = Good
			return true
		}
	}
	memo[p] = Bad
	return false
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {}
