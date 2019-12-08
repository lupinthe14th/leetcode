package findthesmallestdivisorgivenathreshold

func smallestDivisor(nums []int, threshold int) int {
	const MaxUint = ^uint(0)
	const MaxInt = int(MaxUint >> 1)
	lo, hi := 0, MaxInt

	for lo+1 < hi {
		m := (lo + hi) >> 1
		var sum int
		for _, n := range nums {
			// Each result of division is rounded to the nearest integer greater than or equal to that element.
			sum += (n + m - 1) / m
		}
		if sum <= threshold {
			hi = m
		} else {
			lo = m
		}
	}
	return hi
}
