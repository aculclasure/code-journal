package brackets

import (
	"errors"
)

type stack struct {
	items []rune
}

func (s *stack) push(r rune) {
	s.items = append(s.items, r)
}

func (s *stack) pop() (rune, error) {
	numItems := len(s.items)

	if numItems == 0 {
		return 0, errors.New("stack is empty")
	}

	item := s.items[numItems-1]
	s.items = s.items[:numItems-1]
	return item, nil
}

func (s *stack) isEmpty() bool {
	return len(s.items) == 0
}

func newStack() *stack {
	return &stack{items: make([]rune, 0)}
}

func isALeftBracket(r rune) bool {
	return r == '[' || r == '{' || r == '('
}

var bracketPairs = map[rune]rune{
	'}': '{',
	')': '(',
	']': '[',
}

// Bracket accepts a string and returns true if all pairs of bracket characters
// (e.g '{', '(', '[' ) are matched and returns false otherwise.
func Bracket(input string) bool {
	bracketStack := newStack()

	for _, nextChar := range input {
		if isALeftBracket(nextChar) {
			bracketStack.push(nextChar)
		} else if expectedLeftBracket, ok := bracketPairs[nextChar]; ok {
			poppedLeftBracket, err := bracketStack.pop()
			if err != nil || poppedLeftBracket != expectedLeftBracket {
				return false
			}
		}
	}
	if !bracketStack.isEmpty() {
		return false
	}
	return true
}
