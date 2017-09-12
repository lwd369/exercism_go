package secret

import (
	"strconv"
)

const testVersion = 2

// Handshake make a handshake
func Handshake(code uint) []string {
	binNum := Reverse(parseBinary(code))
	result := []string{}
	for index, char := range binNum {
		if char == rune('1') {
			switch {
			case index == 0:
				result = append(result, "wink")
			case index == 1:
				result = append(result, "double blink")
			case index == 2:
				result = append(result, "close your eyes")
			case index == 3:
				result = append(result, "jump")
			case index == 4:
				result = ReverseArray(result)
			}
		}
	}
	return result
}

// parse a uint number to binary uint number
func parseBinary(code uint) string {
	return strconv.FormatUint(uint64(code), 2)
}

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// ReverseArray reverse a array
func ReverseArray(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
