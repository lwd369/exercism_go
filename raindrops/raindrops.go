package raindrops

import "fmt"

const testVersion = 3

// Convert a num to raindrops string
func Convert(num int) string {
	raindrops := ""
	if num%3 == 0 {
		raindrops += "Pling"
	}
	if num%5 == 0 {
		raindrops += "Plang"
	}
	if num%7 == 0 {
		raindrops += "Plong"
	}

	if len(raindrops) == 0 {
		raindrops += fmt.Sprintf("%d", num)
	}

	return raindrops
}
