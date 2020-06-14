package railfence

import "strings"

func createCipher(numItems, numRails int) []int {
	var rails [][]int
	rail, slope := 0, 1

	for i := 0; i < numItems; i++ {
		if rail <= numRails {
			rails = append(rails, []int{})
		}
		rails[rail] = append(rails[rail], i)
		if rail+slope < 0 || numRails <= rail+slope {
			slope = -slope
		}
		rail += slope
	}
	result := []int{}
	for _, slice := range rails {
		result = append(result, slice...)
	}
	return result
}

func Encode(plainText string, numRails int) string {
	textLength := len(plainText)
	cipheredText := make([]string, textLength)

	for to, from := range createCipher(textLength, numRails) {
		cipheredText[to] = string(plainText[from])
	}
	return strings.Join(cipheredText, "")
}

func Decode(cipheredText string, numRails int) string {
	textLength := len(cipheredText)
	plainText := make([]string, textLength)

	for from, to := range createCipher(textLength, numRails) {
		plainText[to] = string(cipheredText[from])
	}
	return strings.Join(plainText, "")
}
