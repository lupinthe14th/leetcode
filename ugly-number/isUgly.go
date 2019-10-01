package isugly

import "log"

func isUgly(num int) bool {
	for num%2 == 0 {
		num = num / 2
	}
	log.Print(num)
	if num == 1 {
		return true
	}
	for num%3 == 0 {
		num = num / 3
	}
	log.Print(num)
	if num == 1 {
		return true
	}
	for num%5 == 0 {
		num = num / 5
	}
	log.Print(num)
	if num == 1 {
		return true
	}
	return false
}
