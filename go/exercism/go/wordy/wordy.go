package wordy

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func plus(a, b int) int {
	return a + b
}

func minus(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) int {
	return a / b
}

type op func(a, b int) int

var ops = map[string]op{
	"plus":          plus,
	"minus":         minus,
	"multiplied by": multiply,
	"divided by":    divide,
}

// Answer computes the answer to a math question asked in plain English and returns
// a tuple (int, bool) where the int is the answer and the bool indicates whether
// the answer was able to be computed ok.
func Answer(question string) (int, bool) {
	allowedOps := strings.Join(getAllowedOps(), "|")
	regexPattern := fmt.Sprintf(`^What is (-?\d+)((?: (?:%s) -?\d+)*)\?$`, allowedOps)
	match := regexp.MustCompile(regexPattern).FindStringSubmatch(question)
	if len(match) != 3 {
		return 0, false
	}

	answer, _ := strconv.Atoi(match[1])
	tokenPattern := fmt.Sprintf(` (%s) (-?\d+)`, allowedOps)
	tokens := regexp.MustCompile(tokenPattern).FindAllStringSubmatch(match[2], -1)
	for _, t := range tokens {
		operator, _ := ops[t[1]]
		num, _ := strconv.Atoi(t[2])
		answer = operator(answer, num)
	}
	return answer, true
}

func getAllowedOps() []string {
	allowedOps := make([]string, 0, len(ops))

	for k := range ops {
		allowedOps = append(allowedOps, k)
	}
	return allowedOps
}
