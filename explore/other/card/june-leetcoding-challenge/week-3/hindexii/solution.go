package hindexii

// https://leetcode.com/problems/h-index-ii/discuss/71063/Standard-binary-search
func hIndex(citations []int) int {
	n := len(citations)
	lo, hi := 0, n-1
	for lo <= hi {
		mi := (lo + hi) / 2
		if citations[mi] >= n-mi {
			hi = mi - 1
		} else {
			lo = mi + 1
		}
	}
	return n - lo
}
