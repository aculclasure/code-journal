package phonenumber

import (
	"fmt"
	"regexp"
	"strings"
)

// Number accepts a phone number and returns it's
// NANP encoding
func Number(input string) (string, error) {
	input = strings.TrimSpace(input)
	var areaCode, exchange string

	if err := validateInput(input); err != nil {
		return "", err
	}

	digits := getDigits(input)
	switch {
	case len(digits) == 11 && strings.Index(digits, "1") != 0:
		return "", fmt.Errorf("country code can only be 1")
	case len(digits) == 11:
		areaCode = digits[1:4]
		exchange = digits[4:7]
		digits = digits[1:]
	case len(digits) == 10:
		areaCode = digits[0:3]
		exchange = digits[3:6]
	default:
		return "", fmt.Errorf("number of digits in input should only be 10 or 11")
	}

	re := regexp.MustCompile(`^[2-9].*`)
	switch {
	case re.MatchString(areaCode) == false:
		return "", fmt.Errorf("area code should begin with [2-9]")
	case re.MatchString(exchange) == false:
		return "", fmt.Errorf("exchange code should begin with [2-9]")
	default:
		return digits, nil
	}
}

// Format accepts a phone number input and returns the cleaned
// representation as "(XXX) XXX-XXXX" or returns an error
// if the input is an invalid phone number.
func Format(input string) (string, error) {
	number, err := Number(input)

	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s", number[0:3], number[3:6], number[6:]), nil
}

// AreaCode accepts a phone number input and returns the area
// code as a string or returns an error if the input is an
// invalid phone number.
func AreaCode(input string) (string, error) {
	number, err := Number(input)

	if err != nil {
		return "", err
	}
	return number[0:3], nil
}

func validateInput(input string) error {
	re := regexp.MustCompile(`[^\s\+\(\)\d-.]`)

	switch {
	case re.MatchString(input):
		return fmt.Errorf("input should only contain [+,.,-,0-9,whitespace]")
	case strings.Count(input, "+") > 1:
		return fmt.Errorf("input should only contain a single \"+\" preceding the country code")
	case strings.Count(input, "+") == 1 && strings.Index(input, "+") != 0:
		return fmt.Errorf("\"+\" can only be located preceding the country code")
	default:
		return nil
	}
}

func getDigits(input string) string {
	re := regexp.MustCompile(`\d+`)

	return strings.Join(re.FindAllString(input, -1), "")
}
