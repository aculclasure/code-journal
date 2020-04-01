package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	randomGenerator = rand.New(rand.NewSource(time.Now().UnixNano()))
	usedNames       = map[string]bool{}
)

const maxNumUniqueNames = 26 * 26 * 10 * 10 * 10

// Robot represents a robot with a name attribute.
type Robot struct {
	name string
}

// Name returns the name attribute for an existing Robot or generates a new
// name if this is an uninitialized Robot and all possible unique names have
// not been seen. If all possible unique names have already been seen,
// an error is returned.
func (r *Robot) Name() (string, error) {
	if r.name == "" && len(usedNames) < maxNumUniqueNames {
		r.name = getPotentialName()
		for usedNames[r.name] {
			r.name = getPotentialName()
		}
		usedNames[r.name] = true
	} else if r.name == "" {
		return "", fmt.Errorf("all %d possible unique names have been generated already",
			maxNumUniqueNames)
	}
	return r.name, nil
}

// Reset resets the name attribute of a given Robot to the
// default empty value.
func (r *Robot) Reset() {
	r.name = ""
}

func getPotentialName() (name string) {
	firstLetter := string(randomGenerator.Intn(26) + int('A'))
	secondLetter := string(randomGenerator.Intn(26) + int('A'))
	num := randomGenerator.Intn(1000)
	return fmt.Sprintf("%s%s%03d", firstLetter, secondLetter, num)
}
