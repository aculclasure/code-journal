package clock

import (
	"fmt"
)

const minutesPerDay = 1440

// Clock represents a clock
type Clock int

// New accepts an hour and minute and returns a Clock with this time
func New(hour, minute int) Clock {
	totalMinutes := (60*hour + minute) % minutesPerDay

	for totalMinutes < 0 {
		totalMinutes += minutesPerDay
	}
	return Clock(totalMinutes)
}

// String returns a string representation of a given Clock
// (e.g. "08:00" for 8AM)
func (c Clock) String() string {
	hours := c / 60
	minutes := c % 60
	return fmt.Sprintf("%02d:%02d", hours, minutes)
}

// Add adds a given number of minutes to a Clock and returns
// a Clock with the new time
func (c Clock) Add(minutes int) Clock {
	return New(0, int(c)+minutes)
}

// Subtract subtracts a given number of minutes from a Clock and
// returns a Clock with the new time
func (c Clock) Subtract(minutes int) Clock {
	return New(0, int(c)-minutes)
}
