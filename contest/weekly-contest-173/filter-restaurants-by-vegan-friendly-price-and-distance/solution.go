package main

import (
	"sort"
)

// refarence: https://leetcode.com/problems/filter-restaurants-by-vegan-friendly-price-and-distance/discuss/490311/Python-Straight-Forward
// Complexity:
// - Time O(nlogn)
// - Space O(n)
func filterRestaurants(restaurants [][]int, veganFriendly int, maxPrice int, maxDistance int) []int {
	sort.Slice(restaurants, func(i, j int) bool {
		if restaurants[i][1] == restaurants[j][1] {
			return restaurants[i][0] > restaurants[j][0]
		}
		return restaurants[i][1] > restaurants[j][1]
	})

	ans := make([]int, 0)
	for _, restaurant := range restaurants {
		i := restaurant[0]
		v := restaurant[2]
		p := restaurant[3]
		d := restaurant[4]
		if v >= veganFriendly && p <= maxPrice && d <= maxDistance {
			ans = append(ans, i)
		}
	}
	return ans
}

func main() {}
