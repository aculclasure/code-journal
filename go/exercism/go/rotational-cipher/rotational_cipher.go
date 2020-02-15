package rotationalcipher

import (
	"unicode"
)

const modRange = 26

// RotationalCipher takes a plain text string and a shift value and returns
// the rotationally ciphered result of applying the shift value to the
// plain text string.
func RotationalCipher(inputPlain string, inputShiftKey int) string {
	var ciphered string
	var minBoundary int

	if inputShiftKey == 0 {
		return inputPlain
	}

	for _, c := range inputPlain {
		if unicode.IsLetter(c) {
			switch {
			case unicode.IsUpper(c):
				minBoundary = int('A')
			default:
				minBoundary = int('a')
			}
			ciphered += string((int(c)+inputShiftKey-minBoundary)%modRange + minBoundary)
		} else {
			ciphered += string(c)
		}
	}
	return ciphered
}
