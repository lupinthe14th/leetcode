package main

// findBestValue is returned the minimum such integer of the
// Given an integer array arr and a target value target, return the integer value such that when we change all the integers larger than value in the given array to be equal to value, the sum of the array gets as close as possible (in absolute difference) to target.
// In case of a tie, return the minimum such integer.
// Notice that the answer is not neccesarilly a number from arr.
// SeeAlso: - https://leetcode.com/problems/sum-of-mutated-array-closest-to-target/discuss/463222/Java-Binary-search-O(nlogk)-k-is-the-max-value-in-arr
//          - https://leetcode.com/problems/sum-of-mutated-array-closest-to-target/discuss/463268/JavaC++-4ms-binary-search-short-readable-code-+-sorting-solution
// 探している値は1から与えられた配列の最大値の間です。探索を高速化するために二分探索を使います。
//
func findBestValue(arr []int, target int) int {
	max := maxA(arr)

	lo, hi := 0, max

	d := target // diff
	v := max    // prev value

	for lo <= hi {
		mid := (lo + hi) / 2
		sum := sumA(arr, mid)
		cd := abs(target - sum) // current diff

		// current diffとdiffの比較
		// current diffが小さければ最小のdiffを切り替え、探している値vもmidに切り替えます。
		// current diffと同じであれば、vはmidが小さければ切り替えます。
		if cd < d {
			d = cd
			v = mid
		} else if cd == d {
			v = min(v, mid)
		}

		// 合計がターゲットより大きければhiを下げます
		// そうでなければ、loをあげます
		if sum > target {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return v
}

// min is returned the minimum value of x y.
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// abs is returned the absolute value of x.
func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

// sumA is returned the sum of each element with a value greather than N.
func sumA(arr []int, n int) int {
	var sum int
	for _, v := range arr {
		sum += min(v, n)
	}
	return sum
}

// maxA is returned the maximum value of the given array.
func maxA(arr []int) int {
	max := -1 << 31

	for _, v := range arr {
		if max < v {
			max = v
		}
	}
	return max
}

func main() {}
