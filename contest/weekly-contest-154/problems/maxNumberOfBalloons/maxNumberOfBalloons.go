package maxnumberofballoons

import (
	"strings"
)

func maxNumberOfBalloons(text string) int {
	var count int

	chars := strings.Split(strings.ToLower(text), "")
	charsMap := make(map[int]string)

	for i, char := range chars {
		charsMap[i] = char
	}

	balloonMap := initBalloonMap()
	textLength := len(chars)
	i := 0
	for i < textLength {
		for j := 0; j < 7; j++ {
			if balloonMap[j] == charsMap[i] {
				delete(charsMap, i)
				delete(balloonMap, j)
			}
			if len(balloonMap) == 0 {
				count++
				balloonMap = initBalloonMap()
				i = 0
				break
			}
		}
		i++
	}
	return count
}

func initBalloonMap() map[int]string {
	return map[int]string{
		0: "b",
		1: "a",
		2: "l",
		3: "l",
		4: "o",
		5: "o",
		6: "n",
	}
}
