/*Package gotime adds some handy functionality to the Golang time package.

Right now that functionality is limited in scope to parsing ISO-8601 dates
that do not include the full RFC3339 format. For example, something like
time.Time.UnmarshalJSON("2018-07-14") will fail, whereas something like
time.Time.UnmarshalJSON("2018-04-14 08:23:35Z") will succeed. The format
constraints this library uses came from https://en.wikipedia.org/wiki/ISO_8601
on 2018-195T10-05.

The word "format" is used herein to mean "a valid first argument to time.Parse()".

The Get{Date,Time}Format functions interrogate their string parameter, which
should be an ISO-8601-compatible string, to determine the proper parsing
format if possible. The difference between the Fast and Safely functions
is that the former assume the input is valid while the latter will error if not.
The Fast functions are fast because they use as few rudimentary string
functions as possible to narrow down the possibilities and do no verification
of the source value. The Safely functions are safer because they use regular
expressions and determine the type using the "if it quacks like a duck" method,
however they would not prevent you from passing a time like 99:00

Be sure to pass the entire timestamp to each function. The DateParser and TimeParser
variables are set to the Fast functions by default, and this library uses those
variables internally to determine which set of functions to use.
*/
package gotime

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"

	"github.com/juju/errors"
)

// DateParser specifies a function that returns the date format of an ISO timestamp.
var DateParser = GetDateFormatFast

// TimeParser specifies a function that returns the time format of an ISO timestamp.
var TimeParser = GetTimeFormatFast

// Time implements and extends time.Time
type Time struct {
	time.Time
	OriginalFormat string
}

// UnmarshalJSON implements the json.Unmarshaler interface. The time can be in any format.
func (t *Time) UnmarshalJSON(data []byte) error {
	value := strings.Trim(string(data), "\"")
	if value == "" {
		t.Time = time.Time{}
		return nil
	}

	tm, err := Parse(value)
	t.Time = tm

	return err
}

// MarshalJSON implements the json.Marshaler interface. The time will be in OriginalFormat if set.
// If OriginalTime is not set, the function will fall back to time.time.MarshalJSON().
func (t *Time) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return nil, nil
	}
	if t.OriginalFormat != "" {
		return []byte(fmt.Sprintf("\"%s\"", t.Format(t.OriginalFormat))), nil
	}
	return t.Time.MarshalJSON()
}

// UnmarshalText implements the encoding.TextUnmarshaler interface. The time can be in any format.
func (t *Time) UnmarshalText(data []byte) error {
	value := string(data)
	if value == "" {
		t.Time = time.Time{}
		return nil
	}

	tm, err := Parse(value)
	t.Time = tm

	return err
}

// Value returns the embedded time.Time for use with SQL queries
func (t *Time) Value() (driver.Value, error) {
	return t.Time, nil
}

// Scan implements the sql.Scanner interface
func (t *Time) Scan(value interface{}) error {
	var ok bool
	if t.Time, ok = value.(time.Time); !ok {
		return errors.New("Error converting value to time")
	}
	return nil
}
