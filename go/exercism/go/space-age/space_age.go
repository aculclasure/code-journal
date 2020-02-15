package space

// The Planet type refers to the name of the Planet.
type Planet string

const numSecondsInEarthYear float64 = 31557600

var earthAgeConversions = map[Planet]float64{
	"Earth":   1.0,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Age computes and returns an age in Earth years given a number of seconds and a Planet.
func Age(seconds float64, planet Planet) float64 {
	return (seconds / earthAgeConversions[planet]) / numSecondsInEarthYear
}
