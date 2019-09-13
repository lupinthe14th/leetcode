package palindromenumber

import (
	"log"
	"sort"
	"strconv"
	"strings"
)

func isPalindrome(x int) bool {
	slist := strings.Split(strconv.Itoa(x), "")
	// 負の値の場合、回文にはならないのでここで判定
	if slist[0] == "-" {
		return false
	}
	// 整数の場合
	r := reverse(slist)
	if x == r {
		return true
	}
	return false
}

// reverse は、文字列配列から
func reverse(slist []string) int {
	slleng := len(slist)
	type num struct {
		idx int
		num string
	}
	numbers := make([]num, 0, slleng)
	for i := 0; i < slleng; i++ {
		numbers = append(numbers, num{idx: slleng - i, num: slist[i]})
	}
	sort.Slice(numbers, func(i, j int) bool { return numbers[i].idx < numbers[j].idx })
	tmp := make([]string, 0, slleng)
	for _, v := range numbers {
		tmp = append(tmp, v.num)
	}
	r, err := strconv.Atoi(strings.Join(tmp, ""))
	if err != nil {
		log.Fatal(err)
	}
	return r
}
