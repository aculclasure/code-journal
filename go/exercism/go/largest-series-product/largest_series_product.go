package lsproduct

import (
	"fmt"
	"regexp"
)

// LargestSeriesProduct returns the largest product of a specified
// number of consecutive digits.
func LargestSeriesProduct(digits string, span int) (int64, error) {
	var maxProduct int64 = 1

	if span > len(digits) {
		return 0, fmt.Errorf("span must be smaller than string length")
	}
	if span < 1 {
		return 0, fmt.Errorf("span must be greater than 0")
	}
	if !inputIsValid(digits) {
		return 0, fmt.Errorf("digits input must only contain digits")
	}

	for i := 0; i+span <= len(digits); i++ {
		product := getProduct(digits[i : i+span])
		if product > maxProduct {
			maxProduct = product
		}
	}
	return maxProduct, nil
}

func inputIsValid(input string) bool {
	re := regexp.MustCompile(`\D+`)
	return re.MatchString(input) == false
}

func getProduct(digits string) int64 {
	var product int64 = 1

	for _, r := range digits {
		digit := int64(r - '0')
		product *= digit
		if product == 0 {
			return product
		}
	}
	return product
}
