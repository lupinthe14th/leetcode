package averagesalaryexcludingtheminimumandmaximumsalary

func average(salary []int) float64 {
	lo, hi := 1<<31, -1<<31
	sum := 0
	for _, n := range salary {
		sum += n
		lo = min(lo, n)
		hi = max(hi, n)
	}
	return float64(sum-(lo+hi)) / float64(len(salary)-2)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
