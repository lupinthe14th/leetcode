package countnumberofteams

func numTeams(rating []int) int {
	n := len(rating)
	var ans int
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if i != j && j != k && k != i {
					if (rating[i] < rating[j] && rating[j] < rating[k]) || (rating[i] > rating[j] && rating[j] > rating[k]) {
						ans++
					}
				}
			}
		}
	}
	return ans
}
