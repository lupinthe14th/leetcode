package encodenumber

import "strconv"

func encode(num int) string {
	if num == 0 {
		return ""
	}
	num++
	var ans string
	for num > 0 {
		ans = strconv.Itoa((num % 2)) + ans
		num /= 2
	}
	return string([]rune(ans)[1:])
}
