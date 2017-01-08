package match

import (
	"net"
	"strings"
)

const validRunes = "abcdefABCDEF0123456789:"

// TrimAddress will trim non-v4 and v6 characters when using strings.TrimFunc
func TrimAddress(r rune) bool {
	if strings.ContainsRune(validRunes, r) {
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

// MatchIP will return a slice of all IP addresses found in a string
func MatchIP(s string) []string {
	s = stripBrackets(s)
	var matches []string
	// We're going to check each token in the string
	parts := strings.Split(s, " ")
	for _, part := range parts {
		part := strings.TrimFunc(part, TrimAddress)
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
