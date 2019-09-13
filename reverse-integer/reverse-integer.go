package reverseinteger

import (
	"log"
	"strconv"
	"strings"
)

func reverse(x int) int {
	log.Printf("x: %d", x)
	integerList := strings.Split(strconv.Itoa(x), "")
	log.Printf("integerList : %+v", integerList)
	tmp := make(map[int]string, len(integerList))
	log.Printf("tmp: %v", tmp)
	i := len(integerList)
	log.Printf("len: %d", i)
	if x > 0 {
		for i > 0 {
			log.Printf("i: %d", i)
			for j := 0; j > len(integerList); j++ {
				tmp[len(integerList)-i] = integerList[i]
			}
			i--
		}
	}
	log.Printf("after append tmp: %+v", tmp)
	s := strings.Join(values(tmp), "")
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func values(m map[int]string) []string {
	s := []string{}
	for _, v := range m {
		s = append(s, v)
	}
	return s
}
