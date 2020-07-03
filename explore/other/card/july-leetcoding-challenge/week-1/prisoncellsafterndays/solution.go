package prisoncellsafterndays

// https://leetcode.com/problems/prison-cells-after-n-days/discuss/205684/JavaPython-Find-the-Loop-or-Mod-14
func prisonAfterNDays(cells []int, N int) []int {
	seen := make(map[[8]int]int)
	var t [8]int
	for i, n := range cells {
		t[i] = n
	}
	for N > 0 {
		var temp [8]int
		seen[t] = N
		N--
		for j := 1; j < 7; j++ {
			if t[j+1] == t[j-1] {
				temp[j] = 1
			} else {
				temp[j] = 0
			}
		}
		t = temp
		if n, ok := seen[t]; ok {
			N %= n - N
		}
	}

	out := make([]int, 8)
	for i, n := range t {
		out[i] = n
	}
	return out
}
