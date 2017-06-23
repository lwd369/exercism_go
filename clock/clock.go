package clock

import (
	"fmt"
)

const testVersion = 4

// Clock contains hour and minute
type Clock struct {
	hour, minute int
}

// New create a new Clock struct
func New(hour, minute int) Clock {
	if minute >= 60 {
		hour = hour + minute/60
		minute = minute % 60
	} else if minute < 0 {
		hour = hour + minute/60 - 1
		minute = 60 + minute%60
		if minute%60 == 0 {
			minute = 0
			hour++
		}
	}

	if hour >= 24 {
		hour = hour % 24
	} else if hour < 0 {
		hour = (24 + hour%24) % 24
	}

	return Clock{hour, minute}
}

func (c Clock) String() string {
	hour, minute := fmt.Sprint(c.hour), fmt.Sprint(c.minute)
	if c.hour < 10 {
		hour = "0" + hour
	}

	if c.minute < 10 {
		minute = "0" + minute
	}

	return fmt.Sprintf("%s:%s", hour, minute)
}

// Add minutes to clock, and return a new Clock struct
func (c Clock) Add(minutes int) Clock {
	hour := c.hour
	minute := c.minute + minutes
	return New(hour, minute)
}
