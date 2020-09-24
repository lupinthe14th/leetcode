package gasstation

func canCompleteCircuit(gas []int, cost []int) int {

	out := -1
	var dfs func(start, cur, tank int)
	dfs = func(start, cur, tank int) {
		tank -= cost[cur]
		if tank < 0 {
			return
		}
		cur++
		if cur == len(gas) {
			cur = 0
		}
		tank += gas[cur]
		if start == cur {
			out = start
			return
		}
		dfs(start, cur, tank)
	}

	for i, g := range gas {
		dfs(i, i, g)
	}
	return out
}
