package gigasecond

// import path for the time package from the standard library
import (
	"time"
)

const testVersion = 4

// AddGigasecond return time after 100000000seconds
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1e9 * time.Second)
}
