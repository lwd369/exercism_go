package hamming

import (
	"errors"
)

const testVersion = 6

// Distance compare two string's different chars number
func Distance(a, b string) (distance int, err error) {
	if len(a) != len(b) {
		err = errors.New("argument error")
		return
	}

	for index := 0; index < len(a); index++ {
		if a[index] != b[index] {
			distance++
		}
	}
	return
}
