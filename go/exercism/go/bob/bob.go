// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

// Hey analyzes a given remark and returns a response depending on what type of remark it is.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	result := ""
	switch {
	case isQuestion(remark) && isShouting(remark):
		result = "Calm down, I know what I'm doing!"
	case isQuestion(remark) && !isShouting(remark):
		result = "Sure."
	case !isQuestion(remark) && isShouting(remark):
		result = "Whoa, chill out!"
	case isEmpty(remark):
		result = "Fine. Be that way!"
	default:
		result = "Whatever."
	}
	return result
}

func hasLetters(remark string) bool {
	return strings.IndexFunc(remark, unicode.IsLetter) != -1
}

func isQuestion(remark string) bool {
	return strings.HasSuffix(remark, "?")
}

func isExclamation(remark string) bool {
	return strings.HasSuffix(remark, "!")
}

func isShouting(remark string) bool {
	return hasLetters(remark) && strings.ToUpper(remark) == remark
}

func isEmpty(remark string) bool {
	return strings.TrimSpace(remark) == ""
}
