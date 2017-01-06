package v6match

import (
	"regexp"
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

// MatchIPv6 will return a slice of all IPv6 address found in a string
func MatchIPv6(s string) []string {
	// FIXME This can match some funky prefix lengths...
	v46re := regexp.MustCompile(`^(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))(/[0-9]{1,3})?$`)

	// clean up paranoid stuff like 1(:)2(:)3(::)
	clean := strings.NewReplacer(
		"(", "", ")", "",
		"[", "", "]", "",
		"{", "", "}", "",
	)
	s = clean.Replace(s)

	var matches []string

	// We're going to check each token in the string
	parts := strings.Split(s, " ")
	for _, part := range parts {
		part := strings.TrimFunc(part, TrimV6)
		match := v46re.FindString(part)
		if len(match) != 0 {
			matches = append(matches, match)
		}
	}
	return matches
}
