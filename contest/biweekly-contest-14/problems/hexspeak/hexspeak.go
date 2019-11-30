package hexspeak

import (
	"fmt"
	"strconv"
)

func toHexspeak(num string) string {
	n, err := strconv.Atoi(num)
	if err != nil {
		panic(err)
	}
	h := fmt.Sprintf("%X", n)
	h = replaceAll(h, "1", "I")
	h = replaceAll(h, "0", "O")

	c := h
	for _, s := range []string{"A", "B", "C", "D", "E", "F", "I", "O"} {
		c = replaceAll(c, s, "")
	}

	if len(c) == 0 {
		return h
	}
	return "ERROR"
}

func replaceAll(s, o, n string) string {
	var ns string

	for _, v := range s {
		if string(v) == o {
			ns += n
		} else {
			ns += string(v)
		}
	}
	return ns
}
