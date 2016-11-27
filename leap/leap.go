package leap

const testVersion = 2

// IsLeapYear checks whether the passed year is a leap year
func IsLeapYear(year int) bool {
	return (year % 4  == 0) && ((year % 100 != 0) || (year % 400) == 0)
}
