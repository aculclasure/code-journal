// Package acronym contains functions that generate acronyms from a given input phrase.
package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate generates and returns an acronym from a given input phrase.
func Abbreviate(s string) string {
	acronym := ""
	splitResults := strings.FieldsFunc(s, split)

	for _, nextField := range splitResults {
		if len(nextField) > 0 && unicode.IsLetter(rune(nextField[0])) {
			acronym += string(unicode.ToUpper(rune(nextField[0])))
		}
	}

	return acronym
}

func split(r rune) bool {
	return r == ' ' || r == '-' || r == '_'
}
