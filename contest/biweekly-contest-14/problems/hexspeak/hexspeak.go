package hexspeak

import (
	"fmt"
	"strconv"
	"strings"
)

func toHexspeak(num string) string {
	n, err := strconv.Atoi(num)
	if err != nil {
		panic(err)
	}
	h := fmt.Sprintf("%X", n)
	h = strings.ReplaceAll(h, "1", "I")
	h = strings.ReplaceAll(h, "0", "O")

	c := h
	for _, s := range []string{"A", "B", "C", "D", "E", "F", "I", "O"} {
		c = strings.ReplaceAll(c, s, "")
	}

	if len(c) == 0 {
		return h
	}
	return "ERROR"
}
