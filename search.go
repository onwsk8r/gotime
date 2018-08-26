package gotime

import (
	"time"
)

// NthWeekday returns a UTC time.Time representing the nth weekday of a month.
// The month is determined by the year and month parameters, `n` is the
// week number, and day is the day. As this function uses time.Time.Add()
// internally, if you ask for, for example, the 20th Saturday, you'll receive
// a time.Time a few months in the future. The times must be in UTC for calculating
// times around daylight savings time.
func NthWeekday(year int, month time.Month, n int, day time.Weekday) time.Time {
	t := FirstWeekday(year, month, day)

	// time.Duration is an int
	return t.Add(7 * time.Duration(n-1) * 24 * time.Hour)
}

// LastWeekday returns a UTC time.Time representing the last <day> in <month>.
func LastWeekday(year int, month time.Month, day time.Weekday) time.Time {
	t := FirstWeekday(year, month, day)

	for {
		t = t.Add(24 * 7 * time.Hour)
		if t.Day() > 24 { // This will catch 31-day months
			break
		}
		if t.Month() != month { // This will catch the rest
			t = t.Add(-24 * 7 * time.Hour)
			break
		}
	}
	return t
}

// FirstWeekday returns a UTC time.Time representing the first <day> in <month>.
func FirstWeekday(year int, month time.Month, day time.Weekday) time.Time {
	t := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)

	var first time.Weekday
	if t.Weekday() > day {
		first = 8 - t.Weekday() + day
	} else {
		first = 1 - t.Weekday() + day
	}

	return t.Add(time.Duration(first-1) * 24 * time.Hour)
}

// Below is a SQL version of the NthWeekday function
// from which that function was derived

// CREATE OR REPLACE FUNCTION public.nthwday(
// 	n integer,
// 	d integer DEFAULT date_part('dow', CURRENT_DATE),
// 	m integer DEFAULT date_part('month', CURRENT_DATE),
// 	y integer DEFAULT date_part('year', CURRENT_DATE))
//     RETURNS integer
//     LANGUAGE 'plpgsql'

//     COST 1
//     IMMUTABLE PARALLEL SAFE
// AS $BODY$

// DECLARE
//     first_dow integer := extract(DOW FROM make_date(y, m, 1));
//     first_d integer;
// BEGIN
//     -- TODO There has to be cleaner way to do this
//     IF first_dow > d THEN
//         first_d = 8 - first_dow + d;
//     ELSE
//     	first_d = 1 - first_dow + d;
//     END IF;
//     RETURN 7*(n-1)+first_d;
// END

// $BODY$;
