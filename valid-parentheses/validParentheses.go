package validparentheses

import (
	"strings"
)

func isValid(s string) bool {
	brackets := strings.Split(s, "")
	if len(brackets)%2 != 0 {
		return false
	}

	if brackets[0] == ")" || brackets[0] == "}" || brackets[0] == "]" {
		return false
	}

	for i := 0; i < len(brackets)/2; i++ {
		switch {
		case brackets[i] == "(":
			if brackets[i+1] == "}" || brackets[i+1] == "]" {
				return false
			}
			if brackets[i+1] != ")" {
				if brackets[len(brackets)-i-1] != ")" {
					return false
				}
			}
		case brackets[i] == "{":
			if brackets[i+1] == ")" || brackets[i+1] == "}" {
				return false
			}
			if brackets[i+1] != "}" {
				if brackets[len(brackets)-i-1] != "}" {
					return false
				}
			}
		case brackets[i] == "[":
			if brackets[i+1] == ")" || brackets[i+1] == "}" {
				return false
			}
			if brackets[i+1] != "]" {
				if brackets[len(brackets)-i-1] != "]" {
					return false
				}
			}
		}
	}
	return true
}
