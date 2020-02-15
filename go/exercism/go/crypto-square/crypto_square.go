package cryptosquare

import (
	"regexp"
	"strings"
)

// Encode accepts a message and returns its square-encoding.
func Encode(message string) string {
	var encoded string
	normalized := getNormalized(message)
	numRows, numCols := getDimensions(len(normalized))
	codeSquare := getInitCodeSquare(numRows, numCols)
	populateCodeSquare(normalized, &codeSquare)
	encoded = getEncodedMessage(codeSquare)

	return encoded
}

func getDimensions(msgSize int) (rows, cols int) {
	for rows = 1; msgSize-(cols*rows) > 0; rows++ {
		for cols = rows; msgSize-(cols*rows) > 0 && cols+1-rows <= 1; cols++ {
		}
		if msgSize-(cols*rows) <= 0 {
			break
		}
	}
	return
}

func getEncodedMessage(codeSquare [][]rune) string {
	var encodedPhrases []string

	if len(codeSquare) < 1 || len(codeSquare[0]) < 1 {
		return ""
	}

	for c := 0; c < len(codeSquare[0]); c++ {
		nextPhrase := ""
		for r := 0; r < len(codeSquare); r++ {
			if codeSquare[r][c] == 0 {
				nextPhrase += " "
			} else {
				nextPhrase += string(codeSquare[r][c])
			}
		}
		encodedPhrases = append(encodedPhrases, nextPhrase)
	}
	return strings.Join(encodedPhrases, " ")
}

func getInitCodeSquare(rows, cols int) [][]rune {
	codeSquare := make([][]rune, rows)

	for i := range codeSquare {
		codeSquare[i] = make([]rune, cols)
	}
	return codeSquare
}

func getNormalized(input string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	strippedInput := re.ReplaceAllString(input, "")

	return strings.ToLower(strippedInput)
}

func populateCodeSquare(message string, codeSquare *[][]rune) {
	var numCols int

	if len((*codeSquare)) > 0 && len((*codeSquare)[0]) > 0 {
		numCols = len((*codeSquare)[0])
	}
	for i, v := range message {
		nextRow, nextCol := i/numCols, i%numCols
		(*codeSquare)[nextRow][nextCol] = v
	}
}
