// Package leap contains a function used to determine if a given year is a leap year.
package leap

// IsLeapYear analyzes a given year and returns true if that year is a leap year and false otherwise.
func IsLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				return true
			}
			return false
		}
		return true
	}
	return false
}
