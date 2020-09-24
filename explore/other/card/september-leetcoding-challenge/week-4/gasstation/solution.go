package gasstation

func canCompleteCircuit(gas []int, cost []int) int {

	start, tank, total := 0, 0, 0
	for station := 0; station < len(gas); station++ {
		fuel := gas[station] - cost[station]
		tank += fuel
		total += fuel

		if tank < 0 {
			tank = 0
			start = station + 1
		}
	}

	if total < 0 {
		return -1
	}
	return start
}
