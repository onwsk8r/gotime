package holiday

import (
	"time"
)

// Easter returns the date of Western Easter for the given year.
// In Western Christianity, Easter Sunday is the first Sunday after
// the first full moon after the Equinox. Generally the Paschal full
// moon (the first after the Equinox) happens shortly after the
// Equinox, so Easter subsequently happens around the end of March.
func Easter(year ...int) time.Time {
	y := parseYear(year...)

	g, c := y%19, y/100
	h := int(c-(c/4)-((8*c+13)/25)+19*g+15) % 30
	i := h - (h/28)*(1-(h/28)*(29/(h+1))*((21-g)/11))

	var day = i - (int(y+(y/4)+i+2-c+(c/4)) % 7) + 28
	var month time.Month = 3

	if day > 31 {
		month++
		day -= 31
	}
	return time.Date(y, month, day, 0, 0, 0, 0, time.Local)
}

// GoodFriday returns the date of Good Friday for the given year.
func GoodFriday(year ...int) time.Time {
	y := parseYear(year...)
	return Easter(y).Add(-24 * 2 * time.Hour)
}

// The above Go function was adapted from the below SQL function,
// which in turn was adapted from a C# function

// CREATE OR REPLACE FUNCTION public.when_is_easter(
// 	y integer DEFAULT date_part('year', CURRENT_DATE))
//     RETURNS date
//     LANGUAGE 'plpgsql'

//     COST 1
//     IMMUTABLE PARALLEL SAFE
// AS $BODY$

// DECLARE
//   d integer = 0;
//   m integer = 0;
//   g integer = y % 19;
//   c integer = y / 100;

//   h integer;
//   i integer;
// BEGIN
//     h := (c - round(c/4) - round((8*c + 13)/25) + 19*g + 15)::integer % 30;
//     i := h - round(h/28) * round(1-(h/28) * round(29 / (h+1)) * round((21-g)/11));

//     d := i - ((y + round(y/4)+i+2-c+ round(c / 4))::integer % 7) + 28;
//     m := 3;

//     IF d > 31 THEN
//         m := m + 1;
//         d := d - 31;
//     END IF;

//     RETURN make_date(y, m, d);
// END

// $BODY$;
