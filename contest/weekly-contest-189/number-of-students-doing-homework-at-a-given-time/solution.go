package numberofstudentsdoinghomeworkatagiventime

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	c := 0
	l := len(startTime)

	for i := 0; i < l; i++ {
		if startTime[i] <= queryTime && endTime[i] >= queryTime {
			c++
		}
	}
	return c
}
