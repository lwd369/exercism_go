package leap

const testVersion = 3

// IsLeapYear check if the year is a leap year
func IsLeapYear(year int) bool {
	if div := year % 400; div == 0 {
		return true
	} else if div := year % 100; div == 0 {
		return false
	} else if div := year % 4; div == 0 {
		return true
	}
	return false
}
