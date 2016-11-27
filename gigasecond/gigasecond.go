package gigasecond

import "time"

// Constant declaration.
const testVersion = 4
const gigasecond = 1000000000 * time.Second

// AddGigasecond adds 10^9 seconds to the passed time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Duration(gigasecond))
}