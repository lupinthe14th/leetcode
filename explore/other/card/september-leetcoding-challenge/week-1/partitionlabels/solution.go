package partitionlabels

func partitionLabels(S string) []int {
	var last [26]int
	for i := 0; i < len(S); i++ {
		last[S[i]-'a'] = i
	}
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	j, anchor := 0, 0
	out := make([]int, 0)
	for i := 0; i < len(S); i++ {
		j = max(j, last[S[i]-'a'])
		if i == j {
			out = append(out, i-anchor+1)
			anchor = i + 1
		}
	}
	return out
}
