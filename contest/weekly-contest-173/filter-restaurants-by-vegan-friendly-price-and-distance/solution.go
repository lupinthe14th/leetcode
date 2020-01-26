package main

import (
	"sort"
)

func filterRestaurants(restaurants [][]int, veganFriendly int, maxPrice int, maxDistance int) []int {
	type result struct {
		id, rating int
	}

	results := make([]result, 0)
	for _, restaurant := range restaurants {
		id := restaurant[0]
		rating := restaurant[1]
		vegan := restaurant[2]
		price := restaurant[3]
		distance := restaurant[4]
		if veganFriendly == 1 {
			if vegan == veganFriendly && price <= maxPrice && distance <= maxDistance {
				results = append(results, result{id: id, rating: rating})
			}
		} else {
			if price <= maxPrice && distance <= maxDistance {
				results = append(results, result{id: id, rating: rating})
			}
		}
	}
	sort.Slice(results, func(i, j int) bool {
		if results[i].rating == results[j].rating {
			return results[i].id > results[j].id
		}
		return results[i].rating > results[j].rating
	})

	ans := make([]int, len(results))
	for i, v := range results {
		ans[i] = v.id
	}
	return ans
}

func main() {}
