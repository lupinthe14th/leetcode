package implementrand10usingrand7

func rand10() int {
	var a, b, idx int
	for {
		a = rand7()
		b = rand7()
		idx = b + (a-1)*7
		if idx <= 40 {
			return 1 + (idx-1)%10
		}
		a = idx - 40
		b = rand7()

		idx = b + (a-1)*7
		if idx <= 60 {
			return 1 + (idx-1)%10
		}

		a = idx - 60
		b = rand7()

		idx = b + (a-1)*7
		if idx <= 20 {
			return 1 + (idx-1)%10
		}
	}
}

func rand7() int {
	return 0
}
