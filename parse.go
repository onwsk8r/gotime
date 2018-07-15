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

// Parse is like a patternless version of time.Parse.
// It uses the package specified pattern finding functions
// to determine the pattern and then calls time.Parse internally.
func Parse(str string) (time.Time, error) {
	datefmt, err := DateParser(str)
	if err != nil {
		return time.Time{}, errors.Wrap(err, NewParseError(str, "Parsing date format"))
	}
	timefmt, err := TimeParser(str)
	if err != nil {
		return time.Time{}, errors.Wrap(err, NewParseError(str, "Parsing time format"))
	}

	return time.Parse(datefmt+timefmt, str)
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
	} else if s, _ := strconv.Atoi(str); strconv.Itoa(s) == str { // nolint: gas, errcheck
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

// GetTimeFormatFast tries to find the time and TZ format. Fast.
// Being as such it uses fuzzy matching, such as the number of
// colons, to determine the format: it does not ensure the string
// conforms to the format. For example, input of "50:60:70" would
// return "15:04:05", and input of "T99-" would return "20060102".
// Do not use this function if you are unsure that your string contains
// a valid time.
//
// Valid times: HH:MM:SS, HH:MM, HH:MM:SS.nnn, hh, hhmm, hhmmss, hhmmss.nnn,
// any of the previous postfixed with Z or +/-hh(:mm)?
func GetTimeFormatFast(str string) (string, error) {
	timefmt := strings.Builder{}

	// Leave off the date portion
	if strings.Contains(str, "T") {
		str = strings.Split(str, "T")[1]
		timefmt.WriteString("T") // nolint: gas, errcheck
	}

	// Strip off the timezone and nanosecond portions first, if they exist
	var tz string
	for _, ch := range []string{"Z", "+", "-"} {
		if idx := strings.LastIndex(str, ch); idx > 5 {
			str, tz = str[:idx], str[idx:]
			break
		}
	}

	var ns string
	if idx := strings.Index(str, "."); idx != -1 {
		str, ns = str[:idx-1], str[idx-1:]
	}

	// Handle the HMS portion
	if len(str) >= 2 {
		timefmt.WriteString("15") // nolint: gas, errcheck
	}
	switch len(str) {
	case 4:
		timefmt.WriteString("04") // nolint: gas, errcheck
	case 5:
		timefmt.WriteString(":04") // nolint: gas, errcheck
	case 6:
		timefmt.WriteString("0405") // nolint: gas, errcheck
	case 8:
		timefmt.WriteString(":04:05") // nolint: gas, errcheck
	}

	// Handle the NS portion
	if len(ns) > 1 {
		timefmt.WriteString(".") // nolint: gas, errcheck
		ns = ns[1:]
		for range ns {
			timefmt.WriteString("0") // nolint: gas, errcheck
		}
	}

	// Handle the TZ portion
	if len(tz) > 0 {
		timefmt.WriteString(string(tz[0])) // nolint: gas, errcheck
		tz = tz[1:]
		switch len(tz) {
		case 2:
			timefmt.WriteString("07") // nolint: gas, errcheck
		case 4:
			timefmt.WriteString("0700") // nolint: gas, errcheck
		case 5:
			timefmt.WriteString("07:00") // nolint: gas, errcheck
		}
	}

	return timefmt.String(), nil
}

// GetDateFormatSafely uses regexes to parse date formats. *NOT IMPLEMENTED
// Using this function ensures the returned format string will
// work correctly. Use this function when you are unsure if
// your date is properly formatted.
func GetDateFormatSafely(str string) (string, error) {
	return "", errors.NotImplementedf("This function has not been implemented")
}

// GetTimeFormatSafely uses regexes to parse date formats. *NOT IMPLEMENTED
// Using this function ensures the returned format string will
// work correctly. Use this function when you are unsure if
// your time is properly formatted.
func GetTimeFormatSafely(str string) (string, error) {
	return "", errors.NotImplementedf("This function has not been implemented")
}
