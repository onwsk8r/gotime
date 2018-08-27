package holiday

import (
	"time"

	"github.com/onwsk8r/gotime"
)

// NYDay returns the date of New Year's Day.
// If it falls on a weekend, it is observed the following Monday.
func NYDay(year ...int) time.Time {
	y := parseYear(year...)
	return time.Date(y, 1, 1, 0, 0, 0, 0, time.UTC)
}

// InaugurationDay returns the date of the next Inauguration Day.
// Inauguration Day is January 20th, the date the new POTUS is
// swon in. As it only happens every four years, the function will
// return the date of the "next" one; these years can be represented
// as (y % 4 == 1). If 2016 or 2017 is passed as a parameter, for
// example, the function will return Jan 20, 2017. If 2018 is passed,
// the function will return Jan 20, 2021.
func InaugurationDay(year ...int) time.Time {
	y := parseYear(year...)
	for {
		if y%4 == 1 {
			break
		}
		y++
	}
	return time.Date(y, 1, 20, 0, 0, 0, 0, time.UTC)
}

// MLKDay returns the date of Martin Luther King Day.
// MLK Day is the third Monday in January and is a Federal holiday.
func MLKDay(year ...int) time.Time {
	y := parseYear(year...)
	return gotime.NthWeekday(y, time.January, 3, time.Monday)
}

// PresidentsDay returns the date of President's Day.
// President's Day is the third Monday in February and a Federal holiday.
func PresidentsDay(year ...int) time.Time {
	y := parseYear(year...)
	return gotime.NthWeekday(y, time.February, 3, time.Monday)
}

// MemorialDay returns the date of Memorial Day.
// Memorial Day is the last Monday in May and a Federal holiday.
func MemorialDay(year ...int) time.Time {
	y := parseYear(year...)
	return gotime.LastWeekday(y, time.May, time.Monday)
}

// IndependenceDay returns the date of Independence Day.
// Independence Day is aka July 4. It is not a Federal Holiday (just kidding!)
func IndependenceDay(year ...int) time.Time {
	y := parseYear(year...)
	return time.Date(y, time.July, 4, 0, 0, 0, 0, time.UTC)
}

// LaborDay returns the date of Labor Day.
// Labor Day is the first Monday in September and a Federal holiday.
func LaborDay(year ...int) time.Time {
	y := parseYear(year...)
	return gotime.FirstWeekday(y, time.September, time.Monday)
}

// ColumbusDay returns the date of Columbus Day.
// Columbus Day is the second Monday in October and a Federal holiday.
func ColumbusDay(year ...int) time.Time {
	y := parseYear(year...)
	return gotime.NthWeekday(y, time.October, 2, time.Monday)
}

// VeteransDay returns the date of Veteran's Day.
// Veteran's Day is Nov 11, and a Federal holiday.
func VeteransDay(year ...int) time.Time {
	y := parseYear(year...)
	return time.Date(y, time.November, 11, 0, 0, 0, 0, time.UTC)
}

// Thanksgiving returns the date of Thanksgiving.
// Thanksgiving is the fourth Thursday in November and a Federal holiday.
func Thanksgiving(year ...int) time.Time {
	y := parseYear(year...)
	return gotime.NthWeekday(y, time.November, 4, time.Thursday)
}

// BlackFriday returns the date of Black Friday.
// Black Friday is the day after Thanksgiving and might as well be a Federal holiday.
func BlackFriday(year ...int) time.Time {
	return Thanksgiving(year...).Add(24 * time.Hour)
}

// ChristmasDay returns the date of Christmas Day.
// Christmas Day is Dec 25, and a Federal holiday.
func ChristmasDay(year ...int) time.Time {
	y := parseYear(year...)
	return time.Date(y, time.December, 25, 0, 0, 0, 0, time.UTC)
}

// ChristmasEve returns the date of Christmas Eve.
// Christmas Eve is the day before Christmas (Day) and might as well be a Federal holiday.
func ChristmasEve(year ...int) time.Time {
	return ChristmasDay(year...).Add(-24 * time.Hour)
}

// NYEve returns the date of New Year's Eve.
// New Year's Eve is the day before Christmas (Day) and might as well be a Federal holiday.
func NYEve(year ...int) time.Time {
	return NYDay(year...).Add(-24 * time.Hour)
}

// Parse year keeps the functions above DRY
func parseYear(year ...int) (y int) {
	if len(year) == 0 {
		y = time.Now().Year()
	} else {
		y = year[0]
	}
	return // nolint: nakedret
}
