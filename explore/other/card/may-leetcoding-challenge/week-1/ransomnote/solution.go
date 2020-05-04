package ransomnote

import (
	"strings"
)

func canConstruct(ransomNote string, magazine string) bool {
	arr := strings.Split(magazine, "")
	c := 0
	for i := range ransomNote {
		for j := 0; j < len(arr); j++ {
			if string(ransomNote[i]) == arr[j] {
				c++
				arr = append(arr[:j], arr[j+1:]...)
				break
			}
		}
	}
	return len(ransomNote) == c
}
