// Package raindrops contains a function that translates numbers into raindrop sounds.
package raindrops

import (
	"strconv"
)

var sounds = map[int]string{
	3: "Pling",
	5: "Plang",
	7: "Plong",
}

// Convert accepts an integer and translates it into a string representing
// raindrop sounds.
func Convert(num int) string {
	sound := ""
	for _, factor := range []int{3, 5, 7} {
		if num%factor == 0 {
			sound += sounds[factor]
		}
	}
	if sound == "" {
		sound = strconv.Itoa(num)
	}

	return sound
}
