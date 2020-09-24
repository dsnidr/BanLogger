package strdur

// String duration package

import (
	"regexp"
	"strconv"
	"strings"
)

const (
	minRegex   = "^[0-9]{1,5}min$"
	hourRegex  = "^[0-9]{1,5}h$"
	dayRegex   = "^[0-9]{1,5}d$"
	weekRegex  = "^[0-9]{1,5}w$"
	monthRegex = "^[0-9]{1,5}m$"
	yearRegex  = "^[0-9]{1,5}y$"
	permRegex  = "^perm$"

	secondsInMinute = 60
	secondsInHour   = 3600
	secondsInDay    = 86400
	secondsInWeek   = 604800
	secondsInMonth  = 28512000
	secondsInYear   = 31556952
)

// GetDuration takes in a duration string like 1min, 1h, 1d, 1w, perm etc
// and returns it's duration in seconds.
func GetDuration(durationString string) int64 {
	// if string is "perm" return 0
	if durationString == "perm" {
		return 0
	}

	// min check
	match, _ := regexp.MatchString(minRegex, durationString)
	if match {
		return getMinDuration(durationString)
	}

	// h check
	match, _ = regexp.MatchString(hourRegex, durationString)
	if match {
		return getHourDuration(durationString)
	}

	// d check
	match, _ = regexp.MatchString(dayRegex, durationString)
	if match {
		return getDayDuration(durationString)
	}

	// w check
	match, _ = regexp.MatchString(weekRegex, durationString)
	if match {
		return getWeekDuration(durationString)
	}

	// m check
	match, _ = regexp.MatchString(monthRegex, durationString)
	if match {
		return getMonthDuration(durationString)
	}

	// y check
	match, _ = regexp.MatchString(yearRegex, durationString)
	if match {
		return getYearDuration(durationString)
	}

	// If there was no match, return 0
	return 0
}

func getMinDuration(durationString string) int64 {
	return getNumberPortion(durationString, "min") * secondsInMinute
}

func getHourDuration(durationString string) int64 {
	return getNumberPortion(durationString, "h") * secondsInHour
}

func getDayDuration(durationString string) int64 {
	return getNumberPortion(durationString, "d") * secondsInDay
}

func getWeekDuration(durationString string) int64 {
	return getNumberPortion(durationString, "w") * secondsInWeek
}

func getMonthDuration(durationString string) int64 {
	return getNumberPortion(durationString, "m") * secondsInMonth
}

func getYearDuration(durationString string) int64 {
	return getNumberPortion(durationString, "y") * secondsInYear
}

func getNumberPortion(durationString, sep string) int64 {
	numberString := strings.Split(durationString, sep)[0]

	number, err := strconv.Atoi(numberString)
	if err != nil {
		return 0
	}

	return int64(number)
}
