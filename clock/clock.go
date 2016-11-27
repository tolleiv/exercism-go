package clock

import "fmt"

const testVersion = 4

type Clock struct {
	hour int
	minute int
}

// Create a new Clock
func New(hour, minute int) Clock {
	init := Clock{hour: normalizeHours(hour)}
	return init.Add(minute)
}

// Print the time with leading zeros
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

// Add allows to add and substract time from the clock
func (c Clock) Add(minutes int) Clock {
	h, m := normalizeMinutes(c.minute + minutes)
	h = normalizeHours(c.hour + h)
	return Clock{hour: h, minute: m}
}

// Makes sure we stay within the 24h bounds
func normalizeHours(hour int) int {
	return (hour%24 + 24) % 24
}

// Makes sure we stay within the 60m bounds
func normalizeMinutes(minutes int) (int, int) {
	minute := (minutes%60 + 60) % 60
	hour := (minutes-minute)/60
	return hour, minute
}