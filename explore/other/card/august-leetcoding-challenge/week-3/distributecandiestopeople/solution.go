package distributecandiestopeople

func distributeCandies(candies int, num_people int) []int {
	out := make([]int, num_people)

	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	for i := 0; candies > 0; i++ {
		out[i%num_people] += min(candies, i+1)
		candies -= i + 1
	}
	return out
}
