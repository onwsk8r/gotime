/*Package holiday contains functions for working with holidays.

While the majority of holidays occur on either a particular date (eg the 25th)
or on a particular day (eg the fourth Thursday), Easter is calculated from a
combination of integer arithmetic and Divine Intervention, which is why it gets
its own file. Good Friday lives in that file as well.

The package is (currently, mainly) built around us.go, which contains functions to
calculate holidays that are observed in the United States. It contains functions
for calculating the 11 federal holidays: that's the 10 you know plus Inauguration
Day, which is observed only in DC to mitigate traffic. Each of those functions
implements the Finder type.

The List type is an array of Finder functions and has functions to identify a given
time.Time as a holiday or observed holiday. There are premade lists for US federal
holidays (aka bank holidays or "days everyone else gets off") and trading holidays:
days the stock markets are closed. More information about the federal and trading
holidays can be found at <https://www.redcort.com/us-federal-bank-holidays> and
<https://www.nyse.com/markets/hours-calendars>.
*/
package holiday

import (
	"time"

	"github.com/onwsk8r/gotime"
)

// TradingHolidays are days the US stock markets are closed.
// This list does not include the days the markets close early: on July 3,
// the day before Thanksgiving, and Christmas Eve the markets close at 1pm ET.
var TradingHolidays List = []Finder{
	NYDay,
	MLKDay,
	PresidentsDay,
	GoodFriday,
	MemorialDay,
	IndependenceDay,
	LaborDay,
	Thanksgiving,
	ChristmasDay,
}

// FederalHolidays are days the US government takes off.
// The USPS and banks tend to observe these holidays as well. This
// list does not include Inauguration Day, as it is only a holiday
// under very specific circumstances.
var FederalHolidays List = []Finder{
	NYDay,
	MLKDay,
	PresidentsDay,
	MemorialDay,
	IndependenceDay,
	LaborDay,
	ColumbusDay,
	VeteransDay,
	Thanksgiving,
	ChristmasDay,
}

// Finder is an interface for holiday calculation functions. Each function
// should accept an optional integer parameter for the year and return the
// date of the holiday. If a parameter is not specified, the function
// should calculate the date of the holiday for the current year.
// The returned time should be in local time (ie time.Local) with
// zeroes for hours, minutes, seconds, and nanoseconds.
type Finder func(year ...int) time.Time

// List represents a list of holiday Finders
type List []Finder

// Contains checks if <date> exists in the List.
func (l *List) Contains(date time.Time) bool {
	return CheckExact(date, l)
}

// Observes checks if <date> is an observed holiday.
func (l *List) Observes(date time.Time) bool {
	return Check(date, l)
}

// Observed returns the observed date of a holiday.
// Generally if a holiday falls on a Saturday it is observed the
// preceeding Friday, and if it falls on a Sunday it is observed
// the following Monday. The only exception to this rule is New
// Year's Day, which is observed the following Monday when it
// occurs on a Saturday.
func Observed(holiday time.Time) time.Time {
	// NYE Exception
	if holiday.Day() == 1 && holiday.Month() == time.January &&
		holiday.Weekday() == time.Saturday {
		holiday = holiday.Add(48 * time.Hour)
	}
	switch holiday.Weekday() {
	case time.Saturday:
		return holiday.Add(-24 * time.Hour)
	case time.Sunday:
		return holiday.Add(24 * time.Hour)
	}
	return holiday
}

// Check whether the given date is a work holiday
func Check(date time.Time, against *List) bool {
	y := date.Year()

	for _, holiday := range *against {
		if gotime.DateEquals(date, Observed(holiday(y))) {
			return true
		}
	}
	return false
}

// CheckExact whether the given date is an actual holiday.
// This function differs from the Check function in that it does not
// consider observed holidays.
func CheckExact(date time.Time, against *List) bool {
	y := date.Year()

	for _, holiday := range *against {
		if gotime.DateEquals(date, holiday(y)) {
			return true
		}
	}
	return false
}

// The Check function was adapted from the below SQL function

// CREATE OR REPLACE FUNCTION public.is_holiday(
// 	d date DEFAULT CURRENT_DATE)
//     RETURNS boolean
//     LANGUAGE 'plpgsql'

//     COST 1
//     IMMUTABLE LEAKPROOF PARALLEL SAFE
// AS $BODY$

// DECLARE
// 	y integer; -- the current year
//     m integer; -- interim month for calculations
//     hdate date; -- interim date for holiday
// BEGIN
// 	y := EXTRACT(YEAR FROM d);

//     -- New Year's Day, observed the Monday after for weekends
//     hdate := make_date(y, 1, 1);
//     LOOP
//     	EXIT WHEN EXTRACT(DOW FROM hdate) NOT IN (0,6);
//         hdate := hdate + 1;
//     END LOOP;
//     IF d = hdate THEN RETURN true; END IF;

//     -- Martin Luther King Day, 3rd Mon in Jan
//     m := 1;
//     hdate := make_date(y, m, nthwday(3,1,m,y));
//     IF d = hdate THEN RETURN true; END IF;

//     -- President's Day, 3rd Mon in Feb
//     m := 2;
//     hdate := make_date(y, m, nthwday(3,1,m,y));
//     IF d = hdate THEN RETURN true; END IF;

//     -- Good Friday, Good Luck!
//     hdate := when_is_easter(y) - 2;
//     IF d = hdate THEN RETURN true; END IF;

//     -- Memorial Day, last Monday in May
//     m := 5;
//     BEGIN
//     	hdate := make_date(y, m, nthwday(5,1,m,y));
//     	EXCEPTION WHEN datetime_field_overflow THEN
//     		hdate := make_date(y, m, nthwday(4,1,m,y));
//     END;
//     IF d = hdate THEN RETURN true; END IF;

//     -- Independence Day, July 4th
//     hdate := make_date(y, 7, 4);
//     IF d = hdate THEN RETURN true; END IF;

//     -- Labor Day, 1st Mon in Sep
//     m := 9;
//     hdate := make_date(y, m, nthwday(1,1,m,y));
//     IF d = hdate THEN RETURN true; END IF;

//     -- Thanksgiving, 4th Thu in Nov
//     m := 11;
//     hdate := make_date(y, m, nthwday(4,4,m,y));
//     IF d = hdate THEN RETURN true; END IF;

//     -- Christmas, Dec 25th
//     hdate := make_date(y, 12, 25);
//     IF d = hdate THEN RETURN true; END IF;

//     RETURN false;
// END

// $BODY$;
