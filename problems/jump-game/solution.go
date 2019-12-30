package main

// canJump is determined if you are able to reach the last index.
// Algorthmus: recursive backtrack
func canJump(nums []int) bool {
	return canJumpFromPosition(0, nums)
}

func canJumpFromPosition(p int, nums []int) bool {
	l := len(nums)
	// ポジションが最終地点に到達していればtrueを返す
	if p == l-1 {
		return true
	}

	// 最も遠いジャンプ地点を決める
	// スライスにおける現在地点とその値を足した値が最も遠いジャンプ地点。
	// もしそれがスライス長を超えるならスライス長が最も遠いジャンプ地点
	// となるため両者を比較して小さい方を最も遠いジャンプ地点となる。
	fj := min(p+nums[p], l-1) //  furthest Jump

	// 次の地点を逐次再帰的に調べる。
	for np := p + 1; np <= fj; np++ {
		if canJumpFromPosition(np, nums) {
			return true
		}
	}
	return false
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {}
