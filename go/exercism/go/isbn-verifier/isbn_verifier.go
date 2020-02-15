package isbn

import (
	"regexp"
)

const validISBNLength = 10

// IsValidISBN accepts a string argument and returns true if it
// represents a valid ISBN-10 number or false otherwise
func IsValidISBN(isbn string) bool {
	dehyphenated := regexp.MustCompile(`-`).ReplaceAllString(isbn, "")
	nextMultiple := 10
	sum := 0

	switch {
	case regexp.MustCompile(`\d{9}[\dxX]`).MatchString(dehyphenated) == false:
		fallthrough
	case len(dehyphenated) > validISBNLength:
		return false
	}

	for _, c := range dehyphenated {
		if c == 'x' || c == 'X' {
			sum += 10
		} else {
			sum += nextMultiple * int(c-'0')
		}
		nextMultiple--
	}
	return sum%11 == 0
}

HmsLogicMonitorDeviceGroupIds