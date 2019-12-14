package elementappearingmorethan25insortedarray

func findSpecialInteger(arr []int) int {
	seen := make(map[int]int, 0)

	var ans int
	for _, n := range arr {
		seen[n]++
	}

	for n, c := range seen {
		if c > len(arr)/4 {
			ans = n
			break
		}
	}
	return ans
}
