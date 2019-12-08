package groupthepeoplegiventhegroupsizetheybelongto

import "log"

func groupThePeople(groupSizes []int) [][]int {

	groups := make(map[int][]int, 0)
	for i, gs := range groupSizes {
		gn := gs
		if _, group := groups[gn]; !group {
			group := make([]int, gs-1)
			group = append(group, i)
			log.Print(group)
		} else {
			group = append(group, 1)
		}
		groups[gn] = group
		log.Printf("groups: %v", groups)

	}
	return [][]int{}
}
