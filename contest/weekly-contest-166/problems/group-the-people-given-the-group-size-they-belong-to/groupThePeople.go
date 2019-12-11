package groupthepeoplegiventhegroupsizetheybelongto

func groupThePeople(groupSizes []int) [][]int {

	groups := make(map[int][]int, 0)
	for gi, gv := range groupSizes {
		groups[gv] = append(groups[gv], gi)
	}

	result := make([][]int, 0)
	for i := 1; i <= 500; i++ {
		if groups[i] != nil {
			if len(groups[i])-1 < i {
				result = append(result, groups[i])
			} else {
				tmp := make([]int, 0)
				for _, group := range groups[i] {
					tmp = append(tmp, group)
					if len(tmp) == i {
						result = append(result, tmp)
						tmp = []int{}
					}
				}
			}
		}
	}
	return result
}
