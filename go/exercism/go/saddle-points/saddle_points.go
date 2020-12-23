package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

// Pair represents the row and column indices of a saddle point
// in a Matrix.
type Pair [2]int

// Matrix stores a matrix of integers as rows and columns.
type Matrix struct {
	data [][]int
}

// New creates a new Matrix of integers from a string input or returns
// an error otherwise.
func New(data string) (*Matrix, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("got an empty matrix data input (want a string with 1 or more lines of numbers)")
	}

	rows := strings.Split(data, "\n")
	numCols := len(strings.Split(strings.TrimSpace(rows[0]), " "))
	m := make([][]int, len(rows))
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
	return &Matrix{data: m}, nil
}

// Saddle returns the saddle points of the Matrix m as a list of
// coordinate pairs.
func (m *Matrix) Saddle() []Pair {
	var pairs []Pair
	rows := m.rows()

	for r := range rows {
		c := 0
		for j := range rows[r] {
			if rows[r][c] < rows[r][j] {
				c = j
			}
		}

		for x := range rows {
			if r != x && rows[x][c] < rows[r][c] {
				c = -1
				break
			}
		}

		if c >= 0 {
			pairs = append(pairs, Pair{r, c})
		}
	}
	return pairs
}

// rows returns the rows of the given matrix as a slice of integer slices.
func (m *Matrix) rows() [][]int {
	rows := make([][]int, len(m.data))

	for i, row := range m.data {
		rows[i] = append(rows[i], row...)
	}
	return rows
}
