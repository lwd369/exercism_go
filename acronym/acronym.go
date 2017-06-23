package acronym

import (
	"strings"
)

const testVersion = 3

// Abbreviate Convert a phrase to its acronym.
func Abbreviate(str string) string {
	str = strings.Replace(str, "-", " ", -1)
	words := strings.Fields(str)
	result := ""
	for _, word := range words {
		result += string([]byte{word[0]})
	}

	return strings.ToUpper(result)
}
