package findthetownjudge

func findJudge(N int, trust [][]int) int {
	a := make([]int, N+1)
	for i := 0; i < len(trust); i++ {
		a[trust[i][0]]--
		a[trust[i][1]]++
	}
	for i := 1; i <= N; i++ {
		if a[i] == N-1 {
			return i
		}
	}
	return -1
}
