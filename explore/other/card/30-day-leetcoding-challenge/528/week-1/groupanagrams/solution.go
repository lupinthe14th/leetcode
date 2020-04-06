package groupanagrams

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	memo := make(map[string][]string, 0)

	for _, v := range strs {
		k := sortString(v)
		memo[sortString(v)] = append(memo[k], v)
	}
	ans := [][]string{}
	for _, v := range memo {
		ans = append(ans, v)
	}
	return ans
}

func sortString(a string) string {
	var aa []byte
	for i := 0; i < len(a); i++ {
		aa = append(aa, a[i])
	}
	sort.Slice(aa, func(i, j int) bool { return aa[i] < aa[j] })
	return string(aa)
}
