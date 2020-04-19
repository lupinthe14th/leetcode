package displaytableoffoodordersinarestaurant

import (
	"fmt"
	"sort"
	"strconv"
)

func displayTable(orders [][]string) [][]string {
	foodItems := make(map[string]interface{})
	memo := make(map[string]map[string]int)

	for _, order := range orders {
		foodItems[order[2]] = nil
		if food, ok := memo[order[1]]; !ok {
			food = make(map[string]int)
			memo[order[1]] = food
		}
		memo[order[1]][order[2]]++
	}

	header := make([]string, 0, len(foodItems))
	for k := range foodItems {
		header = append(header, k)
	}
	sort.Strings(header)
	header = append([]string{"Table"}, header...)
	ans := make([][]string, 0, len(memo))
	for k, v := range memo {
		table := make([]string, 0, len(header))
		table = append(table, k)
		for i, name := range header {
			if i == 0 {
				continue
			}
			table = append(table, fmt.Sprint(v[name]))
		}
		ans = append(ans, table)
	}
	sort.SliceStable(ans, func(i, j int) bool {
		a, _ := strconv.Atoi(ans[i][0])
		b, _ := strconv.Atoi(ans[j][0])
		return a < b
	})
	ans = append([][]string{header}, ans...)
	return ans
}
