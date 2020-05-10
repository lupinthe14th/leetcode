package findthetownjudge

func findJudge(N int, trust [][]int) int {
	me := make([]bool, N+1)
	you := make([]int, N+1)
	for i := 0; i < len(trust); i++ {
		me[trust[i][0]] = true
		you[trust[i][1]]++
	}
	for i := 1; i <= N; i++ {
		if !me[i] && you[i] == N-1 {
			return i
		}
	}
	return -1
}
