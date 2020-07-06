package say

import (
	"fmt"
	"strings"
)

var ones = []string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var teens = []string{"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen",
	"sixteen", "seventeen", "eighteen", "nineteen"}
var tens = []string{"", "ten", "twenty", "thirty", "forty", "fifty", "sixty", "seventy",
	"eighty", "ninety"}
var powers = []string{"", "thousand", "million", "billion"}

// Say accepts a number from 0 to 999,999,999,999 and returns
// a string of how the number would be spoken and a bool indicating
// if the given number is within the acceptable range.
func Say(n int64) (string, bool) {
	if n == 0 {
		return "zero", true
	}

	if n < 0 || n > 999999999999 {
		return "", false
	}

	powerGroups := make([]int64, 0)
	for n > 0 {
		powerGroups = append(powerGroups, n%1000)
		n /= 1000
	}

	spokenNumber := ""
	for i := len(powerGroups) - 1; i >= 0; i-- {
		if powerGroups[i] != 0 {
			spokenNumber += getPower(powerGroups[i]) + powers[i] + " "
		}
	}
	return strings.TrimSpace(spokenNumber), true
}

func getPower(n int64) string {
	words := ""
	hundred, ten, one := (n%1000)/100, (n%100)/10, n%10

	if hundred > 0 {
		words += ones[hundred] + " hundred "
	}
	switch {
	case ten == 1:
		words += teens[one] + " "
	case one == 0:
		words += tens[ten] + " "
	case ten == 0:
		words += ones[one] + " "
	default:
		words += fmt.Sprintf("%s-%s ", tens[ten], ones[one])
	}
	return words
}
