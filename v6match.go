package v6match

import (
	"net"
	"strings"
)

const v6chars = "abcdefABCDEF0123456789:"

// TrimV6 will trim non-v6 characters when using strings.TrimFunc
func TrimV6(r rune) bool {
	if strings.ContainsRune(v6chars, r) {
		return false
	}
	return true
}

// stripBrackets removes all brackets that are frequently used to obfuscate addresses
func stripBrackets(s string) string {
	clean := strings.NewReplacer(
		"(", "", ")", "",
		"[", "", "]", "",
		"{", "", "}", "",
	)
	return clean.Replace(s)
}

// MatchIPv6 will return a slice of all IPv6 addresses found in a string
func MatchIPv6(s string) []string {
	s = stripBrackets(s)
	var matches []string
	// We're going to check each token in the string
	parts := strings.Split(s, " ")
	for _, part := range parts {
		part := strings.TrimFunc(part, TrimV6)
		var match string
		// Need to treat parsing CIDRs differently from IPs without lengths
		if strings.Contains(part, "/") {
			_, ip, err := net.ParseCIDR(part)
			if err != nil {
				continue
			}
			match = ip.String()
		} else {
			ip := net.ParseIP(part)
			if ip == nil {
				continue
			}
			match = ip.String()
		}
		matches = append(matches, match)
	}
	return matches
}
