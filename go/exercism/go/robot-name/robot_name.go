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

const (
	minAllowedLetterValue = int('A')
	validLetterValueRange = int('Z') - minAllowedLetterValue + 1
	maxRandomDigitValue   = 9
	numLettersInRobotName = 2
	numDigitsInRobotName  = 3
)

// Robot represents a robot with a name attribute.
type Robot struct {
	name string
}

// Name returns the name attribute for an existing Robot or generates a new
// name and returns it if this is an uninitialized Robot.
func (r *Robot) Name() (string, error) {
	if r.name == "" {
		for {
			r.name = getPotentialName()
			if _, ok := usedNames[r.name]; !ok {
				usedNames[r.name] = true
				break
			}
		}
	}

	return r.name, nil
}

// Reset resets the name attribute of a given Robot to the
// default empty value.
func (r *Robot) Reset() {
	r.name = ""
}

func getPotentialName() string {
	potentialName := ""

	for i := 0; i < numLettersInRobotName; i++ {
		potentialName += getRandomLetter()
	}
	for i := 0; i < numDigitsInRobotName; i++ {
		potentialName += getRandomDigit()
	}
	return potentialName
}

func getRandomLetter() string {
	return string(rune(randomGenerator.Intn(validLetterValueRange) + minAllowedLetterValue))
}

func getRandomDigit() string {
	return fmt.Sprintf("%d", randomGenerator.Intn(maxRandomDigitValue+1))
}
