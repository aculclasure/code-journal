package romannumerals

import (
	"fmt"
)

type romannumeral struct {
	lettering string
	value     int
}

var romanNumerals = []romannumeral{
	{"M", 1000},
	{"CM", 900},
	{"D", 500},
	{"CD", 400},
	{"C", 100},
	{"XC", 90},
	{"L", 50},
	{"XL", 40},
	{"X", 10},
	{"IX", 9},
	{"V", 5},
	{"IV", 4},
	{"I", 1},
}

// ToRomanNumeral converts an integer argument to a Roman numeral and
// returns this value as a string. If the integer argument is invalid
// an error is returned.
func ToRomanNumeral(num int) (string, error) {
	if num <= 0 || num > 3000 {
		return "", fmt.Errorf("invalid argument: %d", num)
	}

	romanNumeral := ""
	for _, nextNumeral := range romanNumerals {
		for num >= nextNumeral.value {
			romanNumeral += nextNumeral.lettering
			num -= nextNumeral.value
		}
		if num == 0 {
			break
		}
	}
	return romanNumeral, nil
}
