package peoplewhoselistoffavoritecompaniesisnotasubsetofanotherlist

import "strings"

func peopleIndexes(favoriteCompanies [][]string) []int {
	maxl := 0
	for _, f := range favoriteCompanies {
		maxl = max(maxl, len(f))
	}

	ans := make([]int, 0, len(favoriteCompanies))
	for i, cs := range favoriteCompanies {
		if maxl == len(cs) {
			ans = append(ans, i)
		} else {
			flag := true
			for j, bs := range favoriteCompanies {
				if i != j && len(bs) >= len(cs) {
					if subset(bs, cs) {
						flag = false
						break
					}
				}
			}
			if flag {

				ans = append(ans, i)
			}
		}
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func subset(bs, cs []string) bool {
	text := strings.Join(bs, " ")
	for _, t := range cs {
		if !strings.Contains(text, t) {
			return false
		}
	}
	return true
}
