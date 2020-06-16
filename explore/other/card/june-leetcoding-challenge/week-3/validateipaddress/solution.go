package validateipaddress

import (
	"strconv"
	"strings"
)

func validIPAddress(IP string) string {
	var out string
	switch {
	case strings.Contains(IP, "."):
		out = validIPv4(IP)
	case strings.Contains(IP, ":"):
		out = validIPv6(IP)
	}
	if out == "" {
		out = "Neither"
	}
	return out
}

func validIPv4(s string) string {
	ip := strings.Split(s, ".")
	if len(ip) != 4 {
		return ""
	}
	for i := range ip {
		if len(ip[i]) > 1 && ip[i][0] == '0' {
			return ""
		}
		p, err := strconv.ParseUint(ip[i], 10, 8)
		if err != nil {
			return ""
		}
		if p < 0 || 255 < p {
			return ""
		}
	}
	return "IPv4"
}

func validIPv6(s string) string {
	ip := strings.Split(s, ":")
	if len(ip) != 8 {
		return ""
	}
	for i := range ip {
		if len(ip[i]) == 0 {
			return ""
		}
		if len(ip[i]) > 4 {
			return ""
		}
		if _, err := strconv.ParseUint(ip[i], 16, 64); err != nil {
			return ""
		}
	}
	return "IPv6"
}
