package encode

import (
	"strconv"
	"unicode"
)

// RunLengthEncode takes an input string and returns
// the run length-encoded equivalent as a string.
func RunLengthEncode(input string) string {
	var (
		encoded, remaining string
		prevChar           rune
		charCount          int = 1
	)

	if len(input) < 2 {
		return input
	}

	prevChar = []rune(input)[0]
	remaining = input[1:]
	for i, nextChar := range remaining {
		switch {
		case nextChar != prevChar:
			if charCount > 1 {
				encoded += strconv.Itoa(charCount) + string(prevChar)
			} else {
				encoded += string(prevChar)
			}
			if i == len(remaining)-1 {
				encoded += string(nextChar)
			}
			charCount = 1
		default:
			charCount++
			if i == len(remaining)-1 {
				encoded += strconv.Itoa(charCount) + string(nextChar)
			}
		}
		prevChar = nextChar
	}
	return encoded
}

// RunLengthDecode accepts a run-length-encoded string and returns
// the decoded string value.
func RunLengthDecode(input string) string {
	var decoded, charCount string

	for _, nextChar := range input {
		switch {
		case unicode.IsDigit(nextChar):
			charCount += string(nextChar)
		case charCount != "":
			if count, err := strconv.Atoi(charCount); err != nil {
				panic(err)
			} else {
				decoded += expand(count, nextChar)
			}
			charCount = ""
		default:
			decoded += string(nextChar)
		}
	}
	return decoded
}

func expand(count int, c rune) string {
	expanded := ""

	for i := 0; i < count; i++ {
		expanded += string(c)
	}
	return expanded
}
