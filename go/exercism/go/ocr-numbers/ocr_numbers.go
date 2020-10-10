package ocr

import (
	"strings"
)

const (
	zero string = `
 _ 
| |
|_|
   `
	one string = `
   
  |
  |
   `
	two string = `
 _ 
 _|
|_ 
   `
	three string = `
 _ 
 _|
 _|
   `
	four string = `
   
|_|
  |
   `
	five string = `
 _ 
|_ 
 _|
   `
	six string = `
 _ 
|_ 
|_|
   `
	seven string = `
 _ 
  |
  |
   `
	eight string = `
 _ 
|_|
|_|
   `
	nine string = `
 _ 
|_|
 _|
   `
)

var digits = map[string]string{
	zero: "0", one: "1", two: "2", three: "3", four: "4", five: "5", six: "6", seven: "7", eight: "8", nine: "9",
}

func recognizeDigit(digit string) (string, bool) {
	d, ok := digits[digit]
	return d, ok
}

func getLineGroups(str string) [][]string {
	var lineGroups [][]string
	for i, line := range strings.Split(str, "\n") {
		if i%4 == 0 {
			lineGroups = append(lineGroups, []string{})
		}
		lineGroups[i/4] = append(lineGroups[i/4], line)
	}
	return lineGroups
}

func readLine(lines []string) []string {
	var chars []string
	for c := 0; c+3 <= len(lines[0]); c += 3 {
		char := ""
		for l := 0; l < 4; l++ {
			char += "\n" + lines[l][c:c+3]
		}
		chars = append(chars, char)
	}
	return chars
}

// Recognize reads an input string consisting of underscores, pipes, and space
// characters and returns the corresponding number.
func Recognize(input string) []string {
	var digits []string
	for i, line := range getLineGroups(input[1:]) {
		digits = append(digits, "")
		for _, digit := range readLine(line) {
			if d, ok := recognizeDigit(digit); ok {
				digits[i] += d
			} else {
				digits[i] += "?"
			}
		}
	}
	return digits
}
