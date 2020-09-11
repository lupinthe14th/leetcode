package bullsandcows

import (
	"fmt"
)

func getHint(secret string, guess string) string {
	var bulls, cows int
	var numbers [10]int

	for i := 0; i < len(secret); i++ {
		s := secret[i] - '0'
		g := guess[i] - '0'
		if s == g {
			bulls++
		} else {
			if numbers[s] < 0 {
				cows++
			}
			if numbers[g] > 0 {
				cows++
			}
			numbers[s]++
			numbers[g]--
		}
	}
	return fmt.Sprintf("%vA%vB", bulls, cows)
}
