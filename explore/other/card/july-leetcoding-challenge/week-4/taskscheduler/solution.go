package taskscheduler

func leastInterval(tasks []byte, n int) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	l := len(tasks)

	var counter [26]int
	hi, hiCount := 0, 0

	for i := range tasks {
		counter[tasks[i]-'A']++
		cur := counter[tasks[i]-'A']
		if hi == cur {
			hiCount++
		} else if hi < cur {
			hi = cur
			hiCount = 1
		}
	}
	partCount := hi - 1
	partLength := n - (hiCount - 1)
	emptySlots := partCount * partLength
	availableTasks := l - hi*hiCount
	idles := max(0, emptySlots-availableTasks)

	return l + idles
}
