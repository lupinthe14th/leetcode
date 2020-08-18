package distributecandiestopeople

func distributeCandies(candies int, num_people int) []int {
	out := make([]int, num_people)

	g := 1
	for candies > 0 {
		for i := 0; i < num_people; i++ {
			if g < candies {
				out[i] += g
				candies -= g
			} else {
				out[i] += candies
				candies = 0
				break
			}
			g++
		}
	}
	return out
}
