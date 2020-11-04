package queenattack

import (
	"fmt"
	"unicode"
)

type point struct {
	col, row float64
}

var columnValues = map[rune]float64{'a': 1, 'b': 2, 'c': 3, 'd': 4, 'e': 5, 'f': 6, 'g': 7, 'h': 8}

// CanQueenAttack accepts strings w and b representing white and black
// algebraic notational positions of 2 queens and returns true if they
// can attack each other and false if they cannot. An error is returned
// if either of the positions w or b is invalid.
func CanQueenAttack(w, b string) (bool, error) {
	wPoint, err := newPoint(w)
	if err != nil {
		return false, err
	}
	bPoint, err := newPoint(b)
	if err != nil {
		return false, err
	}

	switch {
	case wPoint.col == bPoint.col && wPoint.row == bPoint.row:
		return false, fmt.Errorf("w and b are in the same position (want them to be in different positions): %s",
			w)
	case wPoint.col == bPoint.col:
		fallthrough
	case wPoint.row == bPoint.row:
		return true, nil
	default:
		derivative := (wPoint.row - bPoint.row) / (wPoint.col - bPoint.col)
		if derivative == 1 || derivative == -1 {
			return true, nil
		}
		return false, nil
	}
}

func newPoint(algebraicNotation string) (*point, error) {
	if len(algebraicNotation) != 2 {
		return nil, fmt.Errorf("got an invalid algebraic notation (want a 2-character string): %s",
			algebraicNotation)
	}

	runes := []rune(algebraicNotation)
	col, ok := columnValues[runes[0]]
	if !ok {
		return nil, fmt.Errorf("got an invalid column designation (want a-h): %s", string(runes[0]))
	}
	if !unicode.IsNumber(runes[1]) {
		return nil, fmt.Errorf("got an invalid row designation (want 1-8): %v", string(runes[1]))
	}
	row := float64(int(runes[1] - '0'))
	if row < 1 || row > 8 {
		return nil, fmt.Errorf("got an invalid row designation (want 1-8): %v", string(runes[1]))
	}

	return &point{col: col, row: row}, nil
}
