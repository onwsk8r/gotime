package gotime

import (
	"time"
)

// Equalizer abstracts the Gotime comparison functions
type Equalizer func(a, b time.Time) bool

// DateEquals determines whether the date portion of two Times are equal.
// This function considers two times with the same year and the same day
// of the year to be identical, ignoring time zones. Use DateEqualsTZ if
// you're concerned about time zones.
func DateEquals(a, b time.Time) bool {
	return a.Year() == b.Year() && a.YearDay() == b.YearDay()
}

// DateEqualsTZ extends DateEquals by converting both times to UTC.
// After calling UTC() on both times, this function calls DateEquals
// internally and returns the result.
func DateEqualsTZ(a, b time.Time) bool {
	a, b = a.UTC(), b.UTC()
	return DateEquals(a, b)
}

// TimeEquals determines whether the time portion of two Times are equal.
// This function comapres the hours, minutes, and seconds, ignoring nanoseconds
// and time zones. If those are important to you, use the appropriate function.
func TimeEquals(a, b time.Time) bool {
	return a.Hour() == b.Hour() && a.Minute() == b.Minute() && a.Second() == b.Second()
}

// TimeEqualsTZ extends TimeEquals by converting both times to UTC.
// After calling UTC() on both times, this function calls TimeEquals
// internally and returns the result.
func TimeEqualsTZ(a, b time.Time) bool {
	a, b = a.UTC(), b.UTC()
	return TimeEquals(a, b)
}

// TimeEqualsNS extends TimeEquals by comparing nanoseconds as well.
func TimeEqualsNS(a, b time.Time) bool {
	return TimeEquals(a, b) && a.Nanosecond() == b.Nanosecond()
}

// TimeEqualsNSTZ extends TimeEqualsNS by converting both times to UTC.
// After the convertions, this function returns the value of TimeEqualsNS.
func TimeEqualsNSTZ(a, b time.Time) bool {
	a, b = a.UTC(), b.UTC()
	return TimeEqualsNS(a, b)
}

// SameTime determines whether two times refer to the same time, down to the second.
// This is achieved by comparing the Unix timestamps of each time. For a more accurate
// comparison, consider using time.Time.Equal(), which compares nanoseconds.
func SameTime(a, b time.Time) bool {
	return a.Unix() == b.Unix()
}
