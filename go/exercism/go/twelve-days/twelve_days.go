package twelve

import (
	"fmt"
	"strings"
)

type day struct {
	numberWord string
	gift       string
}

var giftsByDay = []day{
	{},
	{"first", "a Partridge in a Pear Tree."},
	{"second", "two Turtle Doves"},
	{"third", "three French Hens"},
	{"fourth", "four Calling Birds"},
	{"fifth", "five Gold Rings"},
	{"sixth", "six Geese-a-Laying"},
	{"seventh", "seven Swans-a-Swimming"},
	{"eighth", "eight Maids-a-Milking"},
	{"ninth", "nine Ladies Dancing"},
	{"tenth", "ten Lords-a-Leaping"},
	{"eleventh", "eleven Pipers Piping"},
	{"twelfth", "twelve Drummers Drumming"},
}

// Verse accepts a day of Christmas (as an int) and returns the corresponding
// verse as a string.
func Verse(day int) string {
	switch {
	case day == 1:
		return fmt.Sprintf("On the %s day of Christmas my true love gave to me: %s",
			giftsByDay[1].numberWord, giftsByDay[1].gift)
	default:
		verse := fmt.Sprintf("On the %s day of Christmas my true love gave to me: %s, ",
			giftsByDay[day].numberWord, giftsByDay[day].gift)
		for i := day - 1; i > 1; i-- {
			verse += fmt.Sprintf("%s, ", giftsByDay[i].gift)
		}
		verse += fmt.Sprintf("and %s", giftsByDay[1].gift)
		return verse
	}
}

// Song returns the "Twelve Days of Christmas" song as a string.
func Song() string {
	song := ""

	for i := 1; i <= 12; i++ {
		song += Verse(i) + "\n"
	}
	return strings.TrimSpace(song)
}
