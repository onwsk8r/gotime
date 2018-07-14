/*Package gotime adds some handy functionality to the Golang time package.
Right now that functionality is limited in scope to parsing ISO-8601 dates
that do not include the full RFC3339 format. For example, something like
time.Time.UnmarshalJSON("2018-07-14") will fail, whereas something like
time.Time.UnmarshalJSON("2018-04-14 08:23:35Z") will succeed. The format
constraints came from https://en.wikipedia.org/wiki/ISO_8601 a/o 2018-195T10-05

The word "format" is used herein to mean "a valid first argument to time.Parse()".

The Get{Date,Time}Format functions interrogate their string parameter, which
should be an ISO-8601-compatible string, to determine the proper parsing
format if possible. The difference between the Fast and Safely functions
is that the former assume the input is valid while the latter will be certain.
The Fast functions are fast because they use as few rudimentary string
functions as possible to narrow down the possibilities and do no verification
of the source value. The Safely functions are safe because they use regular
expressions and determine the type using the "if it quacks like a duck" method.

Be sure to pass the entire timestamp to each function. The DateParser and TimeParser
variables are set to the Fast functions by default, and this library uses those
variables internally to determine which set of functions to use.
*/
package gotime

// DateParser specifies a function that returns the date format of an ISO timestamp.
var DateParser = GetDateFormatFast
