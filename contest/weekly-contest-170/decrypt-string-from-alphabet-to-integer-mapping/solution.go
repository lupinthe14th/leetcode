package main

var alpabets = map[string]string{
	"1":   "a",
	"2":   "b",
	"3":   "c",
	"4":   "d",
	"5":   "e",
	"6":   "f",
	"7":   "g",
	"8":   "h",
	"9":   "i",
	"10#": "j",
	"11#": "k",
	"12#": "l",
	"13#": "m",
	"14#": "n",
	"15#": "o",
	"16#": "p",
	"17#": "q",
	"18#": "r",
	"19#": "s",
	"20#": "t",
	"21#": "u",
	"22#": "v",
	"23#": "w",
	"24#": "x",
	"25#": "y",
	"26#": "z",
}

func freqAlphabets(s string) string {
	l := len(s)

	var ans string

	for i := 0; i < l; i++ {
		if i+2 >= l {
			ans += alpabets[string(s[i])]
		} else {
			if s[i+2] == '#' {
				tmp := s[i : i+3]
				ans += alpabets[string(tmp)]
				i += 2
			} else {
				ans += alpabets[string(s[i])]
			}
		}
	}
	return ans
}

func main() {}
