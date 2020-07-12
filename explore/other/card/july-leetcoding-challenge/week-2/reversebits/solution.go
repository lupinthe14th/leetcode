package reversebits

import (
	"fmt"
	"strconv"
)

func reverseBits(num uint32) uint32 {
	tmp := make([]rune, 0)
	for _, r := range fmt.Sprintf("%032b", num) {
		tmp = append([]rune{r}, tmp...)
	}
	i, _ := strconv.ParseUint(string(tmp), 2, 32)
	return uint32(i)
}
