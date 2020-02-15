// Package etl contains scrabble data migration functions.
package etl

import (
	"strings"
)

// Transform accepts a map of data in the legacy format and returns
// a map of data in the new format.
func Transform(legacyData map[int][]string) map[string]int {
	transformedData := make(map[string]int)

	for letterScore, letters := range legacyData {
		for _, letter := range letters {
			transformedData[strings.ToLower(letter)] = letterScore
		}
	}
	return transformedData
}
