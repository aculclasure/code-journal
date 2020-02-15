// Package gigasecond contains a function that returns a time in the future.
package gigasecond

import (
	"time"
)

// AddGigasecond adds a gigasecond to the given time and returns the future time.
func AddGigasecond(t time.Time) time.Time {
	gigasecondDuration, _ := time.ParseDuration("1000000000s")
	return t.Add(gigasecondDuration)
}
