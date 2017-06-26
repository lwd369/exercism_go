package pangram

import (
	"strings"
)

const testVersion = 1

const ordA = int('a')
const ordZ = int('z')

// IsPangram determine if a sentence is a pangram.
func IsPangram(s string) bool {
	s = strings.ToLower(s)
	isPangram := true
	for c := ordA; c < ordZ; c++ {
		isPangram = isPangram && strings.Contains(s, string(rune(c)))
	}
	return isPangram
}
