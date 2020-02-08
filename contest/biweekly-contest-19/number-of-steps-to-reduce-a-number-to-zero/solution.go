package numberofstepstoreduceanumbertozero

func numberOfSteps(num int) int {
	var c int
	for num != 0 {
		if num&1 == 0 {
			num /= 2
		} else {
			num--
		}
		c++
	}
	return c
}
