package gotime

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/juju/errors"
)

// ParseError represents an error parsing a time string.
// The string being parsed is stored in the From property
// and a description of what exactly went wrong is stored
// in the Problem property.
// This type is exported for use in tests such as
// `err, ok := err.(gotime.ParseError)`.
type ParseError struct {
	From    string
	Problem string
	errors.Err
}

// NewParseError creates a new ParseError.
// While this function is exported in case this package should
// grow to have other packages inside it, you should probably
// avoid using it in your own code unless you want to confuse people.
// The first parameter is the string that was being parsed and the
// second parameter is a description of what went wrong
// (eg "† is not a valid date")
func NewParseError(from, prob string) *ParseError {
	return &ParseError{From: from, Problem: prob}
}

// Error fulfills the error interface.
// It prints both the Problem and the string that caused it (From).
func (p *ParseError) Error() string {
	return fmt.Sprintf("Error parsing '%s': %s", p.From, p.Problem)
}

// ParseFast extends time.Parse to be a little more flexible.
// It does what it can to handle the various flavors of ISO-8601,
// with some caveats. For example, time.Parse() has no way to handle
// week numbers, so the time would have to be converted to a different
// format, which this function does not do.
func ParseFast(str string) (time.Time, error) {
	datefmt, err := GetDateFormatFast(str)
	if err != nil {
		return time.Time{}, err
	}

	return time.Parse(datefmt, str)
}

// GetDateFormatFast attempts to find the date format. Fast.
// Being as such it uses fuzzy matching, such as the number of
// hyphens, to determine the format: it does not ensure the string
// conforms to the format. For example, input of "200-100-50" would
// return "2006-01-02", and input of "10000000" would return "20060102".
// Do not use this function if you are unsure that your string contains
// a valid date.
// Valid options: --MM-DD, --MMDD, YYYY-MM-DD, YYYY-MM, YYYY, YYYYMMDD
// Currently invalid: YYYY-Www, YYYY-DDD, YYYYDDD
func GetDateFormatFast(str string) (string, error) {
	str = strings.Split(str, "T")[0]
	if strings.Index(str, "--") == 0 { // no year, needs to come first
		// Can be --MM-DD or --MMDD
		switch strings.Count(str, "-") {
		case 2:
			return "--0102", nil
		case 3:
			return "--01-02", nil
		}
	} else if strings.Contains(str, "W") { // Year and week
		// TODO: How to use this format with time.Parse()?
		// Can be YYYY-Www, YYYYWww, YYYY-Www-D, YYYYWwwD
		// switch strings.Count(str, "-") {
		// case 2:
		// 	datefmt = "2006-W01-D"
		// case 1:
		// 	datefmt = "2006-W01"
		// }
		return "", NewParseError(str, "Cannot parse week numbers")
	} else if s, _ := strconv.Atoi(str); strconv.Itoa(s) == str { // nolint: gas
		switch len(str) {
		case 7: // YYYYDDD
			return "", NewParseError(str, "Cannot parse day of year")
		case 8: // YYYYMMDD
			return "20060102", nil
		}
	} else if len(str) == 8 { // YYYY-DDD
		// TODO: How to use this format with time.Parse()?
		// return "2006-002"?
		return "", NewParseError(str, "Cannot parse day of year")
	} else if strings.Contains(str, "-") { // Most likely formats
		// YYYY-MM, YYYY-MM-DD
		switch strings.Count(str, "-") {
		case 2:
			return "2006-01-02", nil
		case 1:
			return "2006-01", nil
		}
	}
	return "", NewParseError(str, "Cannot make heads or tails")
}
