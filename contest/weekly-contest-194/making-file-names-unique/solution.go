package makingfilenamesunique

import (
	"fmt"
	"strings"
)

func getFolderNames(names []string) []string {
	seen := make(map[string]int)

	for i := range names {
		k, ok := seen[names[i]]
		if !ok {
			seen[names[i]]++
		} else {
			for j := 0; j < 5e4; j++ {
				ss := []string{names[i], "(", fmt.Sprintf("%d", k+j), ")"}
				s := strings.Join(ss, "")
				if _, ok := seen[s]; !ok {
					seen[s]++
					seen[names[i]] = k + j
					names[i] = s
					break
				}

			}
		}
	}
	return names
}
