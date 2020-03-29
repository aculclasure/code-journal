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

type Robot struct {
	name string
}

func (r *Robot) Name() (string, error) {
	if r.name == "" {
		r.name = getPotentialName()
		for _, ok := usedNames[r.name]; ok; {
			r.name = getPotentialName()
		}
		usedNames[r.name] = true
	}

	return r.name, nil
}

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
