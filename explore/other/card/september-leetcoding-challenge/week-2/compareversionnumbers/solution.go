package compareversionnumbers

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")
	l := max(len(v1), len(v2))

	for i := 0; i < l; i++ {
		n1, n2 := 0, 0
		if len(v1) > i {
			n1, _ = strconv.Atoi(v1[i])
		}
		if len(v2) > i {
			n2, _ = strconv.Atoi(v2[i])
		}
		if n1 > n2 {
			return 1
		} else if n1 < n2 {
			return -1
		}
	}
	return 0
}
