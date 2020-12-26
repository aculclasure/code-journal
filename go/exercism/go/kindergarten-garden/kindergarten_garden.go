package kindergarten

import (
	"fmt"
	"sort"
	"strings"
)

// plants maps the diagram plant code to a plant name.
var plants = map[rune]string{
	'C': "clover",
	'G': "grass",
	'R': "radishes",
	'V': "violets",
}

// Garden represents the garden containing students' plants.
type Garden struct {
	studentRowPosition map[string]int
	plantRows          [][]rune
}

// NewGarden accepts a diagram representing plant rows and a list of child
// names and returns Garden. If the diagram or list of child names is not
// valid, then an error is returned.
func NewGarden(diagram string, children []string) (*Garden, error) {
	rows := strings.Split(diagram, "\n")
	if len(rows) != 3 {
		return nil, fmt.Errorf("want a diagram with 2 rows of plants and each row on it's own line (got %s)", diagram)
	}

	if len(rows[1]) == 0 || len(rows[2]) == 0 {
		return nil, fmt.Errorf("want each row to be non-empty (got %s)", diagram)
	}

	if len(rows[1]) != len(rows[2]) {
		return nil,
			fmt.Errorf("want each diagram row to be equal length, got %d != %d",
				len(rows[1]), len(rows[2]))

	}

	if len(rows[1])%2 != 0 || len(rows[2])%2 != 0 {
		return nil, fmt.Errorf("want each row to have an even number of plants (got %s)", diagram)
	}

	for _, r := range rows[1:] {
		for _, c := range r {
			if _, ok := plants[c]; !ok {
				return nil,
					fmt.Errorf("invalid diagram, want plant code C, G, R, V, got %s",
						string(c))
			}
		}
	}

	if len(children) == 0 {
		return nil, fmt.Errorf("want a non-empty list of children")
	}

	studentRowPosition := map[string]int{}
	for _, c := range children {
		studentRowPosition[c] = 0
	}
	if len(studentRowPosition) != len(children) {
		return nil, fmt.Errorf("want non-duplicated student names (got %q)", children)
	}

	if len(rows[1])+len(rows[2]) != 4*len(studentRowPosition) {
		return nil,
			fmt.Errorf("invalid diagram, want %d plants, got %d",
				4*len(studentRowPosition), len(rows[1])+len(rows[2]))
	}

	sortedNames := make([]string, len(children))
	copy(sortedNames, children)
	sort.Strings(sortedNames)
	for i, name := range sortedNames {
		studentRowPosition[name] = 2*i + 1
	}
	plantRows := [][]rune{[]rune(rows[1]), []rune(rows[2])}

	return &Garden{studentRowPosition: studentRowPosition, plantRows: plantRows}, nil
}

// Plants accepts a child and returns a list of the child's
// plants and a bool indicating whether the child actually
// has planted anything in the given Garden.
func (g *Garden) Plants(child string) ([]string, bool) {
	rowPosition, ok := g.studentRowPosition[child]
	if !ok {
		return nil, false
	}

	planted := []string{}
	for _, row := range g.plantRows {
		p1, p2 := row[rowPosition-1], row[rowPosition]
		planted = append(planted, plants[p1])
		planted = append(planted, plants[p2])
	}
	return planted, true
}
