package bob

import (
	"strings"
)

const testVersion = 3

// Hey response what's bob said
func Hey(s string) string {
	s = trimVoid(s)
	switch {
	case isEmpty(s):
		return "Fine. Be that way!"
	case isYell(s):
		return "Whoa, chill out!"
	case isQuestion(s):
		return "Sure."
	default:
		return "Whatever."
	}
}

func trimVoid(s string) string {
	s = strings.Trim(s, "\t")
	s = strings.Trim(s, "\n")
	s = strings.Trim(s, "\r")
	s = strings.Trim(s, " ")
	return s
}

func isEmpty(s string) bool {
	return s == ""
}

func isQuestion(s string) bool {
	return strings.HasSuffix(s, "?")
}

func isYell(s string) bool {
	if isOnlyNum(s) {
		return false
	}
	upperS := strings.ToUpper(s)
	return upperS == s
}

func isOnlyNum(s string) bool {
	s = strings.ToLower(s)
	const ordA = int('a')
	const ordZ = int('z')
	for _, val := range s {
		if ordC := int(rune(val)); ordC >= ordA && ordC <= ordZ {
			return false
		}
	}
	return true
}
