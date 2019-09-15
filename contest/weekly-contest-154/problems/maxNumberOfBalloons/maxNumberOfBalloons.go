package maxnumberofballoons

import (
	"log"
	"strings"
)

func maxNumberOfBalloons(text string) int {
	var count int
	chars := strings.Split(strings.ToLower(text), "")
	charsMap := make(map[int]string)

	for i, char := range chars {
		charsMap[i] = char
	}

	for i := 0; i < len(chars); i++ {
		log.Printf("charsMap[%d]: %s", i, charsMap[i])
		if strings.Contains("balloons", charsMap[i]) {
			log.Print("contain")
			delete(charsMap, i)
			if (i+1)%7 == 0 {
				count = count + 1
			}

		}
	}
	return count
}
