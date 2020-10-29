package meetup

import (
	"time"
)

// WeekSchedule represents the week of the month.
type WeekSchedule int

const (
	First WeekSchedule = iota
	Second
	Third
	Fourth
	Last
	Teenth
)

// Day returns the day of the meetup in terms of the WeekSchedule.
func Day(weekSchedule WeekSchedule, weekday time.Weekday, month time.Month, year int) int {
	startDate := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)

	switch weekSchedule {
	case First:
		return getWeekday(startDate, weekday, addDay)
	case Second:
		return getWeekday(startDate.AddDate(0, 0, 7), weekday, addDay)
	case Third:
		return getWeekday(startDate.AddDate(0, 0, 14), weekday, addDay)
	case Fourth:
		return getWeekday(startDate.AddDate(0, 0, 21), weekday, addDay)
	case Last:
		return getWeekday(startDate.AddDate(0, 1, -1), weekday, subtractDay)
	case Teenth:
		return getWeekday(startDate.AddDate(0, 0, 12), weekday, addDay)
	}
	return -1
}

func addDay(t time.Time) time.Time {
	return t.AddDate(0, 0, 1)
}

func subtractDay(t time.Time) time.Time {
	return t.AddDate(0, 0, -1)
}

func getWeekday(start time.Time, weekday time.Weekday, f func(t time.Time) time.Time) int {
	for i := 0; i < 7; i++ {
		if start.Weekday() == weekday {
			break
		}
		start = f(start)
	}
	return start.Day()
}
