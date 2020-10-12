package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

// Matrix stores a matrix of integers as rows and columns.
type Matrix [][]int

// New creates a new Matrix of integers from a string input or returns an error otherwise.
func New(data string) (Matrix, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("got an empty matrix data input (want a string with 1 or more lines of numbers)")
	}

	rows := strings.Split(data, "\n")
	numCols := len(strings.Split(strings.TrimSpace(rows[0]), " "))
	m := make(Matrix, len(rows))
	for i, r := range rows {
		cols := strings.Split(strings.TrimSpace(r), " ")
		if len(cols) != numCols {
			return nil, fmt.Errorf("got a row of length %d (want all rows to have length %d): %s",
				len(cols), numCols, r)
		}
		for _, c := range cols {
			num, err := strconv.Atoi(c)
			if err != nil {
				return nil, err
			}
			m[i] = append(m[i], num)
		}
	}
	return m, nil
}

// Rows returns the rows of the Matrix m as a slice of integer slices.
func (m Matrix) Rows() [][]int {
	rows := make([][]int, len(m))

	for i, row := range m {
		rows[i] = append(rows[i], row...)
	}
	return rows
}

// Cols returns the columns of the Matrix m as a slice of integer slices.
func (m Matrix) Cols() [][]int {
	cols := make([][]int, len(m[0]))

	for _, row := range m {
		for i, col := range row {
			cols[i] = append(cols[i], col)
		}
	}
	return cols
}

// Set sets the value of the matrix M to v at the specified row (r) and specified column (c)
// and returns a bool indicating if the set was successful.
func (m Matrix) Set(r, c, v int) bool {
	if r < 0 || r >= len(m) {
		return false
	}
	if c < 0 || c >= len(m[0]) {
		return false
	}
	m[r][c] = v
	return true
}
